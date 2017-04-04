package order

import (
	"fmt"

	"github.com/crypto-bank/proto/currency"
)

// TypeFromString - Converts from string to Type.
func TypeFromString(str string) (_ Type, err error) {
	switch str {
	case "ask", "sell":
		return Ask, nil
	case "bid", "buy":
		return Bid, nil
	}
	err = fmt.Errorf("invalid order type %q", str)
	return
}

// TotalPrice - Total price of an order.
func (order *Order) TotalPrice() *currency.Volume {
	return &currency.Volume{
		Currency: order.Rate.Currency,
		Amount:   order.Rate.Amount * order.Volume.Amount,
	}
}
