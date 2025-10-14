package currency

import (
	"fmt"
)

const (
	CurrencyUSD = "USD"
	CurrencyEUR = "EUR"
	CurrencyGBP = "GBP"
)

// Helper functions for converting between cents and standard currency format

// StandardCurrencyFormatToCents converts a currency amount to cents
// Example: 50.00 -> 5000
func StandardCurrencyFormatToCents(amount float64) int64 {
	return int64(amount * 100)
}

// CentsToStandardCurrencyFormat converts cents to dollars
// Example: 5000 -> 50.00
func CentsToStandardCurrencyFormat(cents int64) float64 {
	return float64(cents) / 100.0
}

// FormatAmount formats a cent amount as a currency string
// Example: 5000 -> "$50.00"
func FormatAmount(cents int64, currency string) string {
	symbol := "$" // Default to USD
	switch currency {
	case CurrencyEUR:
		symbol = "€"
	case CurrencyGBP:
		symbol = "£"
	}

	return fmt.Sprintf("%s%.2f", symbol, CentsToStandardCurrencyFormat(cents))
}
