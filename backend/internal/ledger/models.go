package ledger

import (
	"digitalwallet/backend/pkg/currency"
	"errors"
)

// Account Types - What kind of account is this?
const (
	AccountTypeUserWallet   = "USER_WALLET"   // Individual user's wallet
	AccountTypeSystemFee    = "SYSTEM_FEE"    // Platform fees/revenue
	AccountTypeExternalBank = "EXTERNAL_BANK" // External bank accounts (liability tracking)
)

// Entry Types - Is money going in or out?
const (
	EntryTypeDebit  = "DEBIT"  // Money leaving the account (negative amount)
	EntryTypeCredit = "CREDIT" // Money entering the account (positive amount)
)

// Transaction Types - What kind of operation is this?
const (
	TransactionTypeTransfer   = "TRANSFER"   // User-to-user transfer
	TransactionTypeDeposit    = "DEPOSIT"    // External funds coming in
	TransactionTypeWithdrawal = "WITHDRAWAL" // Funds going out to external account
	TransactionTypeFee        = "FEE"        // Platform fee charge
)

// Validation errors
var (
	ErrInvalidDebitAmount  = errors.New("debit entries must have negative amounts")
	ErrInvalidCreditAmount = errors.New("credit entries must have positive amounts")
	ErrInvalidEntryType    = errors.New("entry type must be DEBIT or CREDIT")
)

// LedgerEntry represents a single entry in the double-entry ledger
// Every financial transaction creates at least two ledger entries that must balance to zero
type LedgerEntry struct {
	ID              string                 `json:"id"`
	AccountID       string                 `json:"account_id"`         // Which account is affected (wallet_id, system account, etc.)
	AccountType     string                 `json:"account_type"`       // USER_WALLET, SYSTEM_FEE, etc.
	Amount          int64                  `json:"amount"`             // Amount in cents (negative for debit, positive for credit)
	Currency        string                 `json:"currency"`           // USD, EUR, etc.
	EntryType       string                 `json:"entry_type"`         // DEBIT or CREDIT
	TransactionID   string                 `json:"transaction_id"`     // Groups entries belonging to same transaction
	TransactionType string                 `json:"transaction_type"`   // TRANSFER, DEPOSIT, WITHDRAWAL, FEE
	CreatedAt       int64                  `json:"created_at"`         // Unix timestamp
	CreatedBy       string                 `json:"created_by"`         // Service or user that created this entry
	Description     string                 `json:"description"`        // Human-readable description
	Metadata        map[string]interface{} `json:"metadata,omitempty"` // Additional context (optional)
}

// Validate ensures the ledger entry follows double-entry bookkeeping rules
func (e *LedgerEntry) Validate() error {
	// Check entry type is valid
	if e.EntryType != EntryTypeDebit && e.EntryType != EntryTypeCredit {
		return ErrInvalidEntryType
	}

	// Debit entries must have negative amounts
	if e.EntryType == EntryTypeDebit && e.Amount >= 0 {
		return ErrInvalidDebitAmount
	}

	// Credit entries must have positive amounts
	if e.EntryType == EntryTypeCredit && e.Amount <= 0 {
		return ErrInvalidCreditAmount
	}

	return nil
}

// ToDTO converts the entry to a user-friendly format with dollar amounts
func (e *LedgerEntry) ToDTO() *LedgerEntryDTO {
	return &LedgerEntryDTO{
		ID:              e.ID,
		AccountID:       e.AccountID,
		Amount:          currency.CentsToStandardCurrencyFormat(e.Amount),
		Currency:        e.Currency,
		EntryType:       e.EntryType,
		TransactionID:   e.TransactionID,
		TransactionType: e.TransactionType,
		CreatedAt:       e.CreatedAt,
		Description:     e.Description,
	}
}

// LedgerEntryDTO is the API response format with dollar amounts instead of cents
type LedgerEntryDTO struct {
	ID              string  `json:"id"`
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"amount"` // In dollars (e.g., 50.00)
	Currency        string  `json:"currency"`
	EntryType       string  `json:"entry_type"`
	TransactionID   string  `json:"transaction_id"`
	TransactionType string  `json:"transaction_type"`
	CreatedAt       int64   `json:"created_at"`
	Description     string  `json:"description"`
}

// AccountBalance stores the cached balance for an account
// This is a performance optimization - we can always recalculate from ledger_entries
type AccountBalance struct {
	AccountID   string `json:"account_id"`
	AccountType string `json:"account_type"`
	Balance     int64  `json:"balance"` // Balance in cents
	Currency    string `json:"currency"`
	UpdatedAt   int64  `json:"updated_at"`
	LastEntryID string `json:"last_entry_id"` // Last ledger entry applied to this balance
}

// ToDTO converts the balance to a user-friendly format with dollar amounts
func (b *AccountBalance) ToDTO() *AccountBalanceDTO {
	return &AccountBalanceDTO{
		AccountID: b.AccountID,
		Balance:   currency.CentsToStandardCurrencyFormat(b.Balance),
		Currency:  b.Currency,
		UpdatedAt: b.UpdatedAt,
	}
}

// AccountBalanceDTO is the API response format
type AccountBalanceDTO struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"` // In standard format (e.g., 50.00)
	Currency  string  `json:"currency"`
	UpdatedAt int64   `json:"updated_at"`
}
