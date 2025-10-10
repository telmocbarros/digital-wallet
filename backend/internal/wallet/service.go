package wallet

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Create a new wallet
func (s *Service) CreateWallet(userId string) error {
	_, err := s.repo.Create(userId)
	if err != nil {
		return err
	}

	return nil
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

func (s *Service) AddCard(walletID string, card *CardDTO) error {
	if err := s.repo.AddCard(walletID, card); err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveCard(walletID, cardID string) error {
	if err := s.repo.RemoveCard(walletID, cardID); err != nil {
		return err
	}

	return nil
}
