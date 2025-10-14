package ledger

import (
	"digitalwallet/backend/pkg/currency"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

var (
	ErrLedgerEntryNotFound    = errors.New("ledger entry not found")
	ErrAccountBalanceNotFound = errors.New("account balance not found")
	ErrTransactionNotBalanced = errors.New("transaction entries do not sum to zero")
	ErrInsufficientBalance    = errors.New("insufficient balance for this operation")
)

// Repository defines the interface for ledger data operations
type Repository interface {
	// Ledger Entry operations
	CreateEntry(entry *LedgerEntry) error
	CreateEntries(entries []*LedgerEntry) error // Atomic: all or nothing
	GetEntryByID(id string) (*LedgerEntry, error)
	GetEntriesByAccountID(accountID string) ([]*LedgerEntry, error)
	GetEntriesByTransactionID(transactionID string) ([]*LedgerEntry, error)

	// Balance operations
	GetBalance(accountID string) (*AccountBalance, error)
	CreateOrUpdateBalance(accountID, accountType string, amountChange int64, lastEntryID string) error
	CalculateBalanceFromEntries(accountID string) (int64, error)

	// Validation
	VerifyTransactionBalance(transactionID string) error
}

// inMemoryRepository implements Repository using in-memory storage
type inMemoryRepository struct {
	entries  []*LedgerEntry
	balances map[string]*AccountBalance // key: accountID
}

// NewRepository creates a new in-memory ledger repository
func NewRepository() Repository {
	return &inMemoryRepository{
		entries:  []*LedgerEntry{},
		balances: make(map[string]*AccountBalance),
	}
}

// CreateEntry creates a single ledger entry
func (r *inMemoryRepository) CreateEntry(entry *LedgerEntry) error {
	// Validate the entry follows double-entry rules
	if err := entry.Validate(); err != nil {
		log.Printf("Error: Invalid ledger entry: %v", err)
		return err
	}

	// Generate ID if not provided
	if entry.ID == "" {
		entry.ID = uuid.New().String()
	}

	// Set creation time if not provided
	if entry.CreatedAt == 0 {
		entry.CreatedAt = time.Now().Unix()
	}

	// Store the entry
	r.entries = append(r.entries, entry)
	log.Printf("Ledger entry created: %s (account: %s, amount: %d, type: %s, txn: %s)",
		entry.ID, entry.AccountID, entry.Amount, entry.EntryType, entry.TransactionID)

	// Update the account balance
	if err := r.CreateOrUpdateBalance(entry.AccountID, entry.AccountType, entry.Amount, entry.ID); err != nil {
		log.Printf("Error updating balance for account %s: %v", entry.AccountID, err)
		return err
	}

	return nil
}

// CreateEntries creates multiple ledger entries atomically
// This is the primary method for creating transactions (which need multiple entries)
func (r *inMemoryRepository) CreateEntries(entries []*LedgerEntry) error {
	// Validate all entries first
	for i, entry := range entries {
		if err := entry.Validate(); err != nil {
			log.Printf("Error: Entry %d invalid: %v", i, err)
			return err
		}
	}

	// Verify the transaction balances to zero
	var sum int64
	for _, entry := range entries {
		sum += entry.Amount
	}
	if sum != 0 {
		log.Printf("Error: Transaction does not balance. Sum: %d", sum)
		return ErrTransactionNotBalanced
	}

	// Create all entries
	for _, entry := range entries {
		if err := r.CreateEntry(entry); err != nil {
			// In a real database, this would be a transaction rollback
			log.Printf("Error: Failed to create entry in transaction: %v", err)
			return err
		}
	}

	log.Printf("Transaction created successfully with %d entries", len(entries))
	return nil
}

// GetEntryByID retrieves a single ledger entry by ID
func (r *inMemoryRepository) GetEntryByID(id string) (*LedgerEntry, error) {
	for _, entry := range r.entries {
		if entry.ID == id {
			return entry, nil
		}
	}
	log.Printf("Error: Ledger entry not found: %s", id)
	return nil, ErrLedgerEntryNotFound
}

// GetEntriesByAccountID retrieves all ledger entries for an account
// This is useful for generating account statements
func (r *inMemoryRepository) GetEntriesByAccountID(accountID string) ([]*LedgerEntry, error) {
	var accountEntries []*LedgerEntry
	for _, entry := range r.entries {
		if entry.AccountID == accountID {
			accountEntries = append(accountEntries, entry)
		}
	}
	log.Printf("Found %d entries for account %s", len(accountEntries), accountID)
	return accountEntries, nil
}

// GetEntriesByTransactionID retrieves all ledger entries for a transaction
// This shows the complete double-entry for a transaction
func (r *inMemoryRepository) GetEntriesByTransactionID(transactionID string) ([]*LedgerEntry, error) {
	var txnEntries []*LedgerEntry
	for _, entry := range r.entries {
		if entry.TransactionID == transactionID {
			txnEntries = append(txnEntries, entry)
		}
	}
	log.Printf("Found %d entries for transaction %s", len(txnEntries), transactionID)
	return txnEntries, nil
}

// GetBalance retrieves the cached balance for an account
func (r *inMemoryRepository) GetBalance(accountID string) (*AccountBalance, error) {
	balance, exists := r.balances[accountID]
	if !exists {
		log.Printf("Error: Account balance not found: %s", accountID)
		return nil, ErrAccountBalanceNotFound
	}
	return balance, nil
}

// CreateOrUpdateBalance updates the cached balance for an account
func (r *inMemoryRepository) CreateOrUpdateBalance(accountID, accountType string, amountChange int64, lastEntryID string) error {
	balance, exists := r.balances[accountID]

	if !exists {
		// Create new balance
		balance = &AccountBalance{
			AccountID:   accountID,
			AccountType: accountType,
			Balance:     amountChange,
			Currency:    currency.CurrencyUSD, // Default currency
			UpdatedAt:   time.Now().Unix(),
			LastEntryID: lastEntryID,
		}
		r.balances[accountID] = balance
		log.Printf("Account balance created: %s with balance %d", accountID, balance.Balance)
	} else {
		// Update existing balance
		balance.Balance += amountChange
		balance.UpdatedAt = time.Now().Unix()
		balance.LastEntryID = lastEntryID
		log.Printf("Account balance updated: %s, new balance: %d", accountID, balance.Balance)
	}

	return nil
}

// CalculateBalanceFromEntries recalculates an account's balance from all ledger entries
// This is the "source of truth" - the cached balance should always match this
func (r *inMemoryRepository) CalculateBalanceFromEntries(accountID string) (int64, error) {
	var balance int64
	for _, entry := range r.entries {
		if entry.AccountID == accountID {
			balance += entry.Amount
		}
	}
	log.Printf("Calculated balance for account %s: %d", accountID, balance)
	return balance, nil
}

// VerifyTransactionBalance verifies that all entries for a transaction sum to zero
// This is a key integrity check for double-entry bookkeeping
func (r *inMemoryRepository) VerifyTransactionBalance(transactionID string) error {
	entries, err := r.GetEntriesByTransactionID(transactionID)
	if err != nil {
		return err
	}

	var sum int64
	for _, entry := range entries {
		sum += entry.Amount
	}

	if sum != 0 {
		log.Printf("Error: Transaction %s does not balance! Sum: %d", transactionID, sum)
		return ErrTransactionNotBalanced
	}

	log.Printf("Transaction %s verified: balances to zero ✓", transactionID)
	return nil
}
