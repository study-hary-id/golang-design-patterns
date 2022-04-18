package factory

import (
	"errors"
	"fmt"
)

// PaymentMethod is an abstraction of each payment method.
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
		// scenario: Replace DebitCardPM to CreditCardPM.
		return new(CreditCardPM), nil
	default:
		return nil, errors.New(
			fmt.Sprintf("error: payment method %d not recognized", m),
		)
	}
}

/*
Defined available payment method, need to add a new payment method?
Just add below
*/

type CashPM struct{}

// Pay prints the successful message of payment process based on the amount.
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash", amount)
}

// This doesn't used, deprecated based from scenario.
type DebitCardPM struct{}

// Pay prints the successful message of payment process based on the amount.
func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card", amount)
}

// CreditCardPM is a replacement payment method for DebitCardPM.
type CreditCardPM struct{}

// Pay prints the successful message of payment process based on the amount.
func (d *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card (new)", amount)
}
