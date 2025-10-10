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
	AddCard(ID string, card *CardDTO) error
	RemoveCard(walletID, cardId string) error
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
func (r *inMemoryRepository) AddCard(walletID string, card *CardDTO) error {
	var wallet Wallet
	for _, w := range r.wallets {
		if w.ID == walletID {
			wallet = w
			break
		}
	}

	if wallet.ID == "" {
		log.Println("Error: Wallet not found when trying to add a card", walletID)
		return pkg.ErrWalletNotFound
	}

	for _, c := range wallet.Cards {
		if card.CardNumber == c.CardNumber {
			log.Println("Error: Attempting to add an existing card to the wallet", card)
			return pkg.ErrCardAlreadyExists
		}
	}

	newCard := Card{
		ID:         uuid.New().String(),
		CardNumber: card.CardNumber,
		ExpiryDate: card.ExpiryDate,
		CVC:        card.CVC,
		CardHolder: card.CardHolder,
	}

	wallet.Cards = append(wallet.Cards, newCard)
	log.Println("Card added to wallet:")
	return nil
}

// RemoveCard implements Repository.
func (r *inMemoryRepository) RemoveCard(walletId string, cardId string) error {
	panic("unimplemented")
}
