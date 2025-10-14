package ledger

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

// Service handles ledger business logic
type Service struct {
	repo Repository
}

// NewService creates a new ledger service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// TransferRequest represents a request to transfer money between accounts
type TransferRequest struct {
	FromAccountID string
	ToAccountID   string
	Amount        int64  // Amount in cents
	Description   string
	TransactionID string // Optional: can be generated if not provided
}

// DepositRequest represents a request to deposit money into an account
type DepositRequest struct {
	AccountID     string
	Amount        int64  // Amount in cents
	Source        string // e.g., "external_bank", "stripe"
	Description   string
	TransactionID string // Optional
}

// WithdrawalRequest represents a request to withdraw money from an account
type WithdrawalRequest struct {
	AccountID     string
	Amount        int64  // Amount in cents
	Destination   string // e.g., "external_bank"
	Description   string
	TransactionID string // Optional
}

// RecordTransfer creates ledger entries for a transfer between two accounts
// This is the core operation for user-to-user transfers
func (s *Service) RecordTransfer(req *TransferRequest) (string, error) {
	// Validate request
	if req.FromAccountID == "" || req.ToAccountID == "" {
		return "", fmt.Errorf("from and to account IDs are required")
	}
	if req.Amount <= 0 {
		return "", fmt.Errorf("amount must be positive")
	}

	// Check if sender has sufficient balance
	fromBalance, err := s.repo.GetBalance(req.FromAccountID)
	if err != nil {
		// If balance doesn't exist yet, it means balance is 0
		if err != ErrAccountBalanceNotFound {
			return "", fmt.Errorf("error checking balance: %w", err)
		}
		// Balance is 0, cannot transfer
		return "", ErrInsufficientBalance
	}

	if fromBalance.Balance < req.Amount {
		log.Printf("Error: Insufficient balance. Account %s has %d, needs %d",
			req.FromAccountID, fromBalance.Balance, req.Amount)
		return "", ErrInsufficientBalance
	}

	// Generate transaction ID if not provided
	transactionID := req.TransactionID
	if transactionID == "" {
		transactionID = uuid.New().String()
	}

	now := time.Now().Unix()

	// Create ledger entries for the transfer
	entries := []*LedgerEntry{
		// Debit from sender
		{
			ID:              uuid.New().String(),
			AccountID:       req.FromAccountID,
			AccountType:     AccountTypeUserWallet,
			Amount:          -req.Amount, // Negative for debit
			Currency:        CurrencyUSD,
			EntryType:       EntryTypeDebit,
			TransactionID:   transactionID,
			TransactionType: TransactionTypeTransfer,
			CreatedAt:       now,
			CreatedBy:       "ledger-service",
			Description:     fmt.Sprintf("Transfer to %s: %s", req.ToAccountID, req.Description),
		},
		// Credit to receiver
		{
			ID:              uuid.New().String(),
			AccountID:       req.ToAccountID,
			AccountType:     AccountTypeUserWallet,
			Amount:          req.Amount, // Positive for credit
			Currency:        CurrencyUSD,
			EntryType:       EntryTypeCredit,
			TransactionID:   transactionID,
			TransactionType: TransactionTypeTransfer,
			CreatedAt:       now,
			CreatedBy:       "ledger-service",
			Description:     fmt.Sprintf("Transfer from %s: %s", req.FromAccountID, req.Description),
		},
	}

	// Create entries atomically
	if err := s.repo.CreateEntries(entries); err != nil {
		log.Printf("Error creating transfer entries: %v", err)
		return "", err
	}

	log.Printf("Transfer recorded: %s -> %s, amount: %d cents, txn: %s",
		req.FromAccountID, req.ToAccountID, req.Amount, transactionID)

	return transactionID, nil
}

// RecordTransferWithFee creates ledger entries for a transfer with a platform fee
// This demonstrates how to handle multi-entry transactions (3+ entries)
func (s *Service) RecordTransferWithFee(req *TransferRequest, feeAmount int64) (string, error) {
	// Validate request
	if req.FromAccountID == "" || req.ToAccountID == "" {
		return "", fmt.Errorf("from and to account IDs are required")
	}
	if req.Amount <= 0 || feeAmount < 0 {
		return "", fmt.Errorf("invalid amounts")
	}

	totalDebit := req.Amount + feeAmount

	// Check if sender has sufficient balance (for amount + fee)
	fromBalance, err := s.repo.GetBalance(req.FromAccountID)
	if err != nil {
		if err != ErrAccountBalanceNotFound {
			return "", fmt.Errorf("error checking balance: %w", err)
		}
		return "", ErrInsufficientBalance
	}

	if fromBalance.Balance < totalDebit {
		log.Printf("Error: Insufficient balance for transfer + fee. Has %d, needs %d",
			fromBalance.Balance, totalDebit)
		return "", ErrInsufficientBalance
	}

	// Generate transaction ID if not provided
	transactionID := req.TransactionID
	if transactionID == "" {
		transactionID = uuid.New().String()
	}

	now := time.Now().Unix()

	// Create ledger entries
	entries := []*LedgerEntry{
		// Debit from sender (amount + fee)
		{
			ID:              uuid.New().String(),
			AccountID:       req.FromAccountID,
			AccountType:     AccountTypeUserWallet,
			Amount:          -totalDebit, // Negative for debit
			Currency:        CurrencyUSD,
			EntryType:       EntryTypeDebit,
			TransactionID:   transactionID,
			TransactionType: TransactionTypeTransfer,
			CreatedAt:       now,
			CreatedBy:       "ledger-service",
			Description:     fmt.Sprintf("Transfer to %s (incl. fee): %s", req.ToAccountID, req.Description),
		},
		// Credit to receiver (just the amount, no fee)
		{
			ID:              uuid.New().String(),
			AccountID:       req.ToAccountID,
			AccountType:     AccountTypeUserWallet,
			Amount:          req.Amount, // Positive for credit
			Currency:        CurrencyUSD,
			EntryType:       EntryTypeCredit,
			TransactionID:   transactionID,
			TransactionType: TransactionTypeTransfer,
			CreatedAt:       now,
			CreatedBy:       "ledger-service",
			Description:     fmt.Sprintf("Transfer from %s: %s", req.FromAccountID, req.Description),
		},
		// Credit to system fee account
		{
			ID:              uuid.New().String(),
			AccountID:       "system-fee-account", // System account ID
			AccountType:     AccountTypeSystemFee,
			Amount:          feeAmount, // Positive for credit
			Currency:        CurrencyUSD,
			EntryType:       EntryTypeCredit,
			TransactionID:   transactionID,
			TransactionType: TransactionTypeFee,
			CreatedAt:       now,
			CreatedBy:       "ledger-service",
			Description:     fmt.Sprintf("Transfer fee from %s", req.FromAccountID),
		},
	}

	// Verify entries balance before creating
	// -totalDebit + amount + fee = 0 ✓

	// Create entries atomically
	if err := s.repo.CreateEntries(entries); err != nil {
		log.Printf("Error creating transfer with fee entries: %v", err)
		return "", err
	}

	log.Printf("Transfer with fee recorded: %s -> %s, amount: %d, fee: %d, txn: %s",
		req.FromAccountID, req.ToAccountID, req.Amount, feeAmount, transactionID)

	return transactionID, nil
}

// RecordDeposit creates ledger entries for depositing money from an external source
func (s *Service) RecordDeposit(req *DepositRequest) (string, error) {
	// Validate request
	if req.AccountID == "" {
		return "", fmt.Errorf("account ID is required")
	}
	if req.Amount <= 0 {
		return "", fmt.Errorf("amount must be positive")
	}

	// Generate transaction ID if not provided
	transactionID := req.TransactionID
	if transactionID == "" {
		transactionID = uuid.New().String()
	}

	now := time.Now().Unix()

	// Create ledger entries
	entries := []*LedgerEntry{
		// Credit user's wallet
		{
			ID:              uuid.New().String(),
			AccountID:       req.AccountID,
			AccountType:     AccountTypeUserWallet,
			Amount:          req.Amount, // Positive for credit
			Currency:        CurrencyUSD,
			EntryType:       EntryTypeCredit,
			TransactionID:   transactionID,
			TransactionType: TransactionTypeDeposit,
			CreatedAt:       now,
			CreatedBy:       "ledger-service",
			Description:     fmt.Sprintf("Deposit from %s: %s", req.Source, req.Description),
		},
		// Debit external bank account (system tracking)
		{
			ID:              uuid.New().String(),
			AccountID:       "external-bank-pool", // System account for external funds
			AccountType:     AccountTypeExternalBank,
			Amount:          -req.Amount, // Negative for debit
			Currency:        CurrencyUSD,
			EntryType:       EntryTypeDebit,
			TransactionID:   transactionID,
			TransactionType: TransactionTypeDeposit,
			CreatedAt:       now,
			CreatedBy:       "ledger-service",
			Description:     fmt.Sprintf("External deposit to %s", req.AccountID),
		},
	}

	// Create entries atomically
	if err := s.repo.CreateEntries(entries); err != nil {
		log.Printf("Error creating deposit entries: %v", err)
		return "", err
	}

	log.Printf("Deposit recorded: %s, amount: %d cents, source: %s, txn: %s",
		req.AccountID, req.Amount, req.Source, transactionID)

	return transactionID, nil
}

// RecordWithdrawal creates ledger entries for withdrawing money to an external account
func (s *Service) RecordWithdrawal(req *WithdrawalRequest) (string, error) {
	// Validate request
	if req.AccountID == "" {
		return "", fmt.Errorf("account ID is required")
	}
	if req.Amount <= 0 {
		return "", fmt.Errorf("amount must be positive")
	}

	// Check if user has sufficient balance
	balance, err := s.repo.GetBalance(req.AccountID)
	if err != nil {
		if err != ErrAccountBalanceNotFound {
			return "", fmt.Errorf("error checking balance: %w", err)
		}
		return "", ErrInsufficientBalance
	}

	if balance.Balance < req.Amount {
		log.Printf("Error: Insufficient balance for withdrawal. Has %d, needs %d",
			balance.Balance, req.Amount)
		return "", ErrInsufficientBalance
	}

	// Generate transaction ID if not provided
	transactionID := req.TransactionID
	if transactionID == "" {
		transactionID = uuid.New().String()
	}

	now := time.Now().Unix()

	// Create ledger entries
	entries := []*LedgerEntry{
		// Debit user's wallet
		{
			ID:              uuid.New().String(),
			AccountID:       req.AccountID,
			AccountType:     AccountTypeUserWallet,
			Amount:          -req.Amount, // Negative for debit
			Currency:        CurrencyUSD,
			EntryType:       EntryTypeDebit,
			TransactionID:   transactionID,
			TransactionType: TransactionTypeWithdrawal,
			CreatedAt:       now,
			CreatedBy:       "ledger-service",
			Description:     fmt.Sprintf("Withdrawal to %s: %s", req.Destination, req.Description),
		},
		// Credit external bank account (system tracking)
		{
			ID:              uuid.New().String(),
			AccountID:       "external-bank-pool",
			AccountType:     AccountTypeExternalBank,
			Amount:          req.Amount, // Positive for credit
			Currency:        CurrencyUSD,
			EntryType:       EntryTypeCredit,
			TransactionID:   transactionID,
			TransactionType: TransactionTypeWithdrawal,
			CreatedAt:       now,
			CreatedBy:       "ledger-service",
			Description:     fmt.Sprintf("External withdrawal from %s", req.AccountID),
		},
	}

	// Create entries atomically
	if err := s.repo.CreateEntries(entries); err != nil {
		log.Printf("Error creating withdrawal entries: %v", err)
		return "", err
	}

	log.Printf("Withdrawal recorded: %s, amount: %d cents, destination: %s, txn: %s",
		req.AccountID, req.Amount, req.Destination, transactionID)

	return transactionID, nil
}

// GetBalance retrieves the current balance for an account
func (s *Service) GetBalance(accountID string) (*AccountBalance, error) {
	balance, err := s.repo.GetBalance(accountID)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// GetAccountStatement retrieves all ledger entries for an account
// This is like a bank statement showing all transactions
func (s *Service) GetAccountStatement(accountID string) ([]*LedgerEntry, error) {
	entries, err := s.repo.GetEntriesByAccountID(accountID)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

// GetTransactionDetails retrieves all ledger entries for a transaction
// This shows the complete double-entry breakdown
func (s *Service) GetTransactionDetails(transactionID string) ([]*LedgerEntry, error) {
	entries, err := s.repo.GetEntriesByTransactionID(transactionID)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

// VerifyAccountBalance verifies that the cached balance matches the calculated balance
// This is important for integrity checks and reconciliation
func (s *Service) VerifyAccountBalance(accountID string) (bool, error) {
	// Get cached balance
	cachedBalance, err := s.repo.GetBalance(accountID)
	if err != nil {
		return false, fmt.Errorf("error getting cached balance: %w", err)
	}

	// Calculate actual balance from entries
	calculatedBalance, err := s.repo.CalculateBalanceFromEntries(accountID)
	if err != nil {
		return false, fmt.Errorf("error calculating balance: %w", err)
	}

	// Compare
	if cachedBalance.Balance != calculatedBalance {
		log.Printf("WARNING: Balance mismatch for account %s! Cached: %d, Calculated: %d",
			accountID, cachedBalance.Balance, calculatedBalance)
		return false, nil
	}

	log.Printf("Balance verification passed for account %s: %d cents", accountID, cachedBalance.Balance)
	return true, nil
}

// VerifyTransaction verifies that a transaction's entries balance to zero
func (s *Service) VerifyTransaction(transactionID string) error {
	return s.repo.VerifyTransactionBalance(transactionID)
}
