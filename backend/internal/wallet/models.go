package wallet

type Card struct {
	ID         string `json:"id"`
	CardNumber string `json:"card_number"`
	EntityID   string `json:"entity"`
	CardHolder string `json:"card_holder"`
	CVC        string `json:"cvc"`
	ExpiryDate int64  `json:"expiry_date"`
	CreatedAt  int64  `json:"created_at"`
}

type CardDTO struct {
	CardNumber string `json:"card_number"`
	Entity     string `json:"entity"`
	CardHolder string `json:"card_holder"`
	CVC        string `json:"cvc"`
	ExpiryDate string `json:"expiry_date"`
}

type Wallet struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
	Cards     []Card `json:"cards"`
}
