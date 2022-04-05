package factory

import (
	"errors"
	"fmt"
)

// PaymentMethod is an abstraction layer of Factory Method.
type PaymentMethod interface {
	Pay(amount float32) string
}

// Enumeration of available payment method.
const (
	Cash      = 1
	DebitCard = 2
)

// GetPaymentMethod returns PaymentMethod instance based on the enumeration.
// If the PaymentMethod isn't available it returns an error.
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(
			fmt.Sprintf("error: payment method %d not recognized\n", m),
		)
	}
}

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

type DebitCardPM struct{}

func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}
