package currency

import "fmt"

// FromString - Creates a currency from a symbol string.
func FromString(symbol string) *Currency {
	return &Currency{Symbol: symbol}
}

// NewPair - Creates a new currency pair from symbols.
func NewPair(first, second string) *Pair {
	return &Pair{
		First:  FromString(first),
		Second: FromString(second),
	}
}

// NewVolume - Creates currency volume structure.
func NewVolume(currency *Currency, amount float64) *Volume {
	return &Volume{
		Currency: currency,
		Amount:   amount,
	}
}

// Concat - Concats currency pair with a symbol.
func (pair *Pair) Concat(symbol string) string {
	return fmt.Sprintf("%s%s%s", pair.First.Symbol, symbol, pair.Second.Symbol)
}

// HumanString - Volume as human-readable string.
func (volume *Volume) HumanString() string {
	return fmt.Sprintf("%.8f %s", volume.Amount, volume.Currency.Symbol)
}
