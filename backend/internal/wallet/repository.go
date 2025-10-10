package wallet

import (
	"digitalwallet/backend/pkg"
	"log"

	"github.com/google/uuid"

	"time"
)

type Repository interface {
	GetByID(ID string) (*Wallet, error)
	GetByUserID(userID string) (*Wallet, error)
	Create(userID string) (string, error)
	AddCard(ID string, card *CardDTO) (string, error)
	RemoveCard(walletID, cardId string) error
	GetCard(walletID, cardID string) (*Card, error)
}

// inMemoryRepository implements Repository using in-memory storage
type inMemoryRepository struct {
	wallets []Wallet
}

func NewRepository() Repository {
	return &inMemoryRepository{
		wallets: []Wallet{},
	}
}

var timeFormat = "02-01-2006"

// Create implements Repository.
func (r *inMemoryRepository) Create(userID string) (string, error) {
	for _, wallet := range r.wallets {
		if wallet.UserID == userID {
			log.Println("Error: User already has a wallet", userID)
			return "", pkg.ErrUserALreadyHasAWallet
		}
	}

	newWallet := Wallet{
		ID:        uuid.New().String(),
		UserID:    userID,
		CreatedAt: time.Now().Unix(),
		Cards:     []Card{},
	}

	r.wallets = append(r.wallets, newWallet)
	log.Println("Wallet created:", newWallet)
	return newWallet.ID, nil
}

// GetByID implements Repository.
func (r *inMemoryRepository) GetByID(ID string) (*Wallet, error) {
	for _, wallet := range r.wallets {
		if wallet.ID == ID {
			return &wallet, nil
		}
	}

	log.Println("Error: Wallet not found", ID)
	return nil, pkg.ErrWalletNotFound
}

// GetByUserID implements Repository.
func (r *inMemoryRepository) GetByUserID(userID string) (*Wallet, error) {
	for _, wallet := range r.wallets {
		if wallet.UserID == userID {
			return &wallet, nil
		}
	}
	log.Println("Error: Wallet not found for user", userID)
	return nil, pkg.ErrWalletNotFound
}

// AddCard implements Repository.
func (r *inMemoryRepository) AddCard(walletID string, card *CardDTO) (string, error) {
	// Find wallet by index to modify the actual slice element
	// IMPORTANT: We must use the index approach here instead of taking the address
	// of the loop variable (&w) because Go reuses the loop variable in each iteration.
	// Taking &w would give us a pointer to the temporary loop variable, not the
	// actual wallet in the slice, causing modifications to be lost.
	walletIndex := -1
	for i := range r.wallets {
		if r.wallets[i].ID == walletID {
			walletIndex = i
			break
		}
	}

	if walletIndex == -1 {
		log.Println("Error: Wallet not found when trying to add a card", walletID)
		return "", pkg.ErrWalletNotFound
	}

	// Get pointer to the actual slice element so modifications persist
	wallet := &r.wallets[walletIndex]

	for _, c := range wallet.Cards {
		if card.CardNumber == c.CardNumber {
			log.Println("Error: Attempting to add an existing card to the wallet", card)
			return "", pkg.ErrCardAlreadyExists
		}
	}

	expiryDate, err := time.Parse(timeFormat, card.ExpiryDate)
	if err != nil {
		log.Println("Error: invalid expiry date format", card.ExpiryDate, err)
		return "", err
	}

	newCard := Card{
		ID:         uuid.New().String(),
		CardNumber: card.CardNumber,
		ExpiryDate: expiryDate.Unix(),
		CVC:        card.CVC,
		CardHolder: card.CardHolder,
	}

	wallet.Cards = append(wallet.Cards, newCard)
	log.Println("Card added to wallet:", newCard.ID)
	return newCard.ID, nil
}

func (r *inMemoryRepository) GetCard(walletID, cardID string) (*Card, error) {
	var wallet Wallet
	for _, w := range r.wallets {
		if w.ID == walletID {
			wallet = w
			break
		}
	}
	if wallet.ID == "" {
		log.Println("Error: Wallet not found when trying to get a card", walletID)
		return nil, pkg.ErrWalletNotFound
	}

	for _, c := range wallet.Cards {
		if c.ID == cardID {
			return &c, nil
		}
	}

	log.Println("Error: Card not found in wallet", wallet, cardID)
	return nil, pkg.ErrCardNotFound
}

// RemoveCard implements Repository.
func (r *inMemoryRepository) RemoveCard(walletId string, cardId string) error {
	// Use index to get pointer to actual slice element (see AddCard for explanation)
	var wallet *Wallet
	for i, w := range r.wallets {
		if w.ID == walletId {
			wallet = &r.wallets[i]
			break
		}
	}

	if wallet == nil {
		log.Println("Error: Wallet not found when trying to remove a card", walletId)
		return pkg.ErrWalletNotFound
	}

	for i, c := range wallet.Cards {
		if c.ID == cardId {
			// Remove the card from the slice
			wallet.Cards = append(wallet.Cards[:i], wallet.Cards[i+1:]...)
			log.Println("Card removed from wallet:", cardId)
			return nil
		}
	}

	log.Println("Error: Card not found in wallet when trying to remove it", cardId)
	return pkg.ErrCardNotFound
}
