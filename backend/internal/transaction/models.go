package transaction

const (
	TransactionTypeDeposit    = "deposit"
	TransactionTypeWithdrawal = "withdrawal"
	TransactionTypeTransfer   = "transfer"
)

type Transaction struct {
	ID              string `json:"id"`
	TransactionType string `json:"transaction_type"`
	Amount          int64  `json:"amount"` // Amount in cents
	Currency        string `json:"currency"`
	Description     string `json:"description"`
	CreatedAt       int64  `json:"created_at"`
	UpdatedAt       int64  `json:"updated_at"`
	FromAccountID   string `json:"from_account_id"`
	ToAccountID     string `json:"to_account_id"`
	Status          string `json:"status"` // e.g., "pending", "completed", "failed"
}
