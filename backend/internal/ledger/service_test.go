package ledger

import (
	"digitalwallet/backend/pkg/currency"
	"testing"
)

// TestTransferBasic tests a simple transfer between two accounts
func TestTransferBasic(t *testing.T) {
	// Setup
	repo := NewRepository()
	service := NewService(repo)

	aliceWalletID := "alice-wallet-123"
	bobWalletID := "bob-wallet-456"

	// Step 1: Alice deposits $100 (so she has funds to transfer)
	t.Log("Step 1: Alice deposits $100")
	depositReq := &DepositRequest{
		AccountID:   aliceWalletID,
		Amount:      10000, // $100.00 in cents
		Source:      "external_bank",
		Description: "Initial deposit",
	}
	txnID1, err := service.RecordDeposit(depositReq)
	if err != nil {
		t.Fatalf("Failed to record deposit: %v", err)
	}
	t.Logf("✓ Deposit recorded with transaction ID: %s", txnID1)

	// Verify Alice's balance
	aliceBalance, err := service.GetBalance(aliceWalletID)
	if err != nil {
		t.Fatalf("Failed to get Alice's balance: %v", err)
	}
	if aliceBalance.Balance != 10000 {
		t.Errorf("Expected Alice's balance to be 10000, got %d", aliceBalance.Balance)
	}
	t.Logf("✓ Alice's balance: %s", currency.FormatAmount(aliceBalance.Balance, currency.CurrencyUSD))

	// Step 2: Alice transfers $50 to Bob
	t.Log("\nStep 2: Alice transfers $50 to Bob")
	transferReq := &TransferRequest{
		FromAccountID: aliceWalletID,
		ToAccountID:   bobWalletID,
		Amount:        5000, // $50.00 in cents
		Description:   "Payment for services",
	}
	txnID2, err := service.RecordTransfer(transferReq)
	if err != nil {
		t.Fatalf("Failed to record transfer: %v", err)
	}
	t.Logf("✓ Transfer recorded with transaction ID: %s", txnID2)

	// Step 3: Verify balances
	t.Log("\nStep 3: Verify final balances")

	aliceBalance, err = service.GetBalance(aliceWalletID)
	if err != nil {
		t.Fatalf("Failed to get Alice's balance: %v", err)
	}
	expectedAlice := int64(5000) // $100 - $50 = $50
	if aliceBalance.Balance != expectedAlice {
		t.Errorf("Expected Alice's balance to be %d, got %d", expectedAlice, aliceBalance.Balance)
	}
	t.Logf("✓ Alice's final balance: %s", currency.FormatAmount(aliceBalance.Balance, currency.CurrencyUSD))

	bobBalance, err := service.GetBalance(bobWalletID)
	if err != nil {
		t.Fatalf("Failed to get Bob's balance: %v", err)
	}
	expectedBob := int64(5000) // $0 + $50 = $50
	if bobBalance.Balance != expectedBob {
		t.Errorf("Expected Bob's balance to be %d, got %d", expectedBob, bobBalance.Balance)
	}
	t.Logf("✓ Bob's final balance: %s", currency.FormatAmount(bobBalance.Balance, currency.CurrencyUSD))

	// Step 4: Verify transaction balances to zero
	t.Log("\nStep 4: Verify transaction integrity")

	if err := service.VerifyTransaction(txnID1); err != nil {
		t.Errorf("Deposit transaction failed verification: %v", err)
	}
	t.Logf("✓ Deposit transaction verified (balances to zero)")

	if err := service.VerifyTransaction(txnID2); err != nil {
		t.Errorf("Transfer transaction failed verification: %v", err)
	}
	t.Logf("✓ Transfer transaction verified (balances to zero)")

	// Step 5: Verify Alice's account balance integrity
	t.Log("\nStep 5: Verify cached balance matches calculated balance")
	balanced, err := service.VerifyAccountBalance(aliceWalletID)
	if err != nil {
		t.Fatalf("Failed to verify Alice's account: %v", err)
	}
	if !balanced {
		t.Errorf("Alice's cached balance does not match calculated balance!")
	}
	t.Logf("✓ Alice's account balance integrity verified")
}

// TestTransferWithFee tests a transfer with a platform fee
func TestTransferWithFee(t *testing.T) {
	// Setup
	repo := NewRepository()
	service := NewService(repo)

	aliceWalletID := "alice-wallet-789"
	bobWalletID := "bob-wallet-012"

	// Step 1: Alice deposits $100
	t.Log("Step 1: Alice deposits $100")
	depositReq := &DepositRequest{
		AccountID:   aliceWalletID,
		Amount:      10000,
		Source:      "stripe",
		Description: "Initial deposit",
	}
	_, err := service.RecordDeposit(depositReq)
	if err != nil {
		t.Fatalf("Failed to record deposit: %v", err)
	}

	// Step 2: Alice transfers $50 to Bob with $1 fee
	t.Log("\nStep 2: Alice transfers $50 to Bob (with $1 fee)")
	transferReq := &TransferRequest{
		FromAccountID: aliceWalletID,
		ToAccountID:   bobWalletID,
		Amount:        5000, // $50.00
		Description:   "Payment with fee",
	}
	feeAmount := int64(100) // $1.00 fee
	txnID, err := service.RecordTransferWithFee(transferReq, feeAmount)
	if err != nil {
		t.Fatalf("Failed to record transfer with fee: %v", err)
	}
	t.Logf("✓ Transfer with fee recorded: %s", txnID)

	// Step 3: Verify balances
	t.Log("\nStep 3: Verify balances")

	aliceBalance, _ := service.GetBalance(aliceWalletID)
	expectedAlice := int64(4900) // $100 - $50 - $1 = $49
	if aliceBalance.Balance != expectedAlice {
		t.Errorf("Expected Alice's balance to be %d, got %d", expectedAlice, aliceBalance.Balance)
	}
	t.Logf("✓ Alice: %s (paid $50 + $1 fee)", currency.FormatAmount(aliceBalance.Balance, currency.CurrencyUSD))

	bobBalance, _ := service.GetBalance(bobWalletID)
	expectedBob := int64(5000) // $50 (no fee)
	if bobBalance.Balance != expectedBob {
		t.Errorf("Expected Bob's balance to be %d, got %d", expectedBob, bobBalance.Balance)
	}
	t.Logf("✓ Bob: %s (received $50, no fee)", currency.FormatAmount(bobBalance.Balance, currency.CurrencyUSD))

	systemBalance, _ := service.GetBalance("system-fee-account")
	expectedSystem := int64(100) // $1 fee
	if systemBalance.Balance != expectedSystem {
		t.Errorf("Expected system balance to be %d, got %d", expectedSystem, systemBalance.Balance)
	}
	t.Logf("✓ System: %s (platform fee)", currency.FormatAmount(systemBalance.Balance, currency.CurrencyUSD))

	// Step 4: Verify transaction balances
	t.Log("\nStep 4: Verify transaction integrity")
	if err := service.VerifyTransaction(txnID); err != nil {
		t.Errorf("Transaction failed verification: %v", err)
	}
	t.Logf("✓ Transaction verified (3 entries balance to zero)")

	// Step 5: Get transaction details
	t.Log("\nStep 5: Get transaction details")
	entries, err := service.GetTransactionDetails(txnID)
	if err != nil {
		t.Fatalf("Failed to get transaction details: %v", err)
	}
	if len(entries) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(entries))
	}
	t.Logf("✓ Transaction has %d entries:", len(entries))
	for _, entry := range entries {
		t.Logf("  - %s %s: %s (%s)",
			entry.AccountID,
			entry.EntryType,
			currency.FormatAmount(entry.Amount, entry.Currency),
			entry.Description,
		)
	}
}

// TestInsufficientBalance tests that transfers fail with insufficient funds
func TestInsufficientBalance(t *testing.T) {
	// Setup
	repo := NewRepository()
	service := NewService(repo)

	aliceWalletID := "alice-wallet-poor"
	bobWalletID := "bob-wallet-lucky"

	// Try to transfer $50 when Alice has $0
	t.Log("Attempting transfer with $0 balance...")
	transferReq := &TransferRequest{
		FromAccountID: aliceWalletID,
		ToAccountID:   bobWalletID,
		Amount:        5000,
		Description:   "This should fail",
	}
	_, err := service.RecordTransfer(transferReq)
	if err != ErrInsufficientBalance {
		t.Errorf("Expected ErrInsufficientBalance, got: %v", err)
	}
	t.Logf("✓ Transfer correctly rejected: %v", err)

	// Give Alice $30, try to transfer $50
	t.Log("\nDeposit $30, then try to transfer $50...")
	depositReq := &DepositRequest{
		AccountID:   aliceWalletID,
		Amount:      3000, // $30
		Source:      "bank",
		Description: "Small deposit",
	}
	service.RecordDeposit(depositReq)

	_, err = service.RecordTransfer(transferReq)
	if err != ErrInsufficientBalance {
		t.Errorf("Expected ErrInsufficientBalance, got: %v", err)
	}
	t.Logf("✓ Transfer correctly rejected: %v", err)

	// Alice should still have $30
	balance, _ := service.GetBalance(aliceWalletID)
	if balance.Balance != 3000 {
		t.Errorf("Expected balance to remain 3000, got %d", balance.Balance)
	}
	t.Logf("✓ Alice's balance unchanged: %s", currency.FormatAmount(balance.Balance, currency.CurrencyUSD))
}

// TestAccountStatement tests retrieving transaction history
func TestAccountStatement(t *testing.T) {
	// Setup
	repo := NewRepository()
	service := NewService(repo)

	aliceWalletID := "alice-wallet-active"

	// Create several transactions
	t.Log("Creating multiple transactions...")

	// Deposit $100
	service.RecordDeposit(&DepositRequest{
		AccountID:   aliceWalletID,
		Amount:      10000,
		Source:      "bank",
		Description: "Salary",
	})

	// Transfer $20 out
	service.RecordTransfer(&TransferRequest{
		FromAccountID: aliceWalletID,
		ToAccountID:   "bob-wallet",
		Amount:        2000,
		Description:   "Lunch payment",
	})

	// Withdraw $30
	service.RecordWithdrawal(&WithdrawalRequest{
		AccountID:   aliceWalletID,
		Amount:      3000,
		Destination: "external_bank",
		Description: "ATM withdrawal",
	})

	// Get account statement
	t.Log("\nRetrieving account statement...")
	entries, err := service.GetAccountStatement(aliceWalletID)
	if err != nil {
		t.Fatalf("Failed to get statement: %v", err)
	}

	t.Logf("✓ Alice's account has %d ledger entries:", len(entries))
	for i, entry := range entries {
		t.Logf("  %d. %s %s: %s - %s",
			i+1,
			entry.EntryType,
			entry.TransactionType,
			currency.FormatAmount(entry.Amount, entry.Currency),
			entry.Description,
		)
	}

	// Verify final balance
	balance, _ := service.GetBalance(aliceWalletID)
	expectedBalance := int64(5000) // $100 - $20 - $30 = $50
	if balance.Balance != expectedBalance {
		t.Errorf("Expected final balance %d, got %d", expectedBalance, balance.Balance)
	}
	t.Logf("✓ Final balance: %s", currency.FormatAmount(balance.Balance, currency.CurrencyUSD))
}
