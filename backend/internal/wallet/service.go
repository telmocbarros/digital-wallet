package wallet

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Create a new wallet
func (s *Service) CreateWallet(userId string) (string, error) {
	walletID, err := s.repo.Create(userId)
	if err != nil {
		return "", err
	}

	return walletID, nil
}

func (s *Service) GetWalletByID(walletId string) (*Wallet, error) {
	wallet, err := s.repo.GetByID(walletId)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (s *Service) GetWalletByUserID(userID string) (*Wallet, error) {
	wallet, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (s *Service) AddCard(walletID string, card *CardDTO) (string, error) {
	cardId, err := s.repo.AddCard(walletID, card)
	if err != nil {
		return "", err
	}

	return cardId, nil
}

func (s *Service) GetCard(walletID, cardID string) (*Card, error) {
	card, err := s.repo.GetCard(walletID, cardID)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (s *Service) RemoveCard(walletID, cardID string) error {
	if err := s.repo.RemoveCard(walletID, cardID); err != nil {
		return err
	}

	return nil
}
