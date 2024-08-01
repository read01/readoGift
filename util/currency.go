package util

const (
	USD = "USD"
	EUR = "EUR"
	CHY = "CHY"
) 


func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD,EUR,CHY :
		return true
	}
	return false
}