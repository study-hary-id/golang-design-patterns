package factory

import (
	"strings"
	"testing"
)

func TestGetAvailablePaymentMethod(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Errorf(
			"GetPaymentMethod(Cash) = %v, %s, nil",
			payment,
			err.Error(),
		)
	}

	payment, err = GetPaymentMethod(DebitCard)
	if err != nil {
		t.Errorf(
			"GetPaymentMethod(DebitCard) = %v, %s, nil",
			payment,
			err.Error(),
		)
	}
}

func TestNotExistsPaymentMethod(t *testing.T) {
	payment, err := GetPaymentMethod(3)
	if err == nil {
		t.Errorf(
			"GetPaymentMethod(3) = %v, want match for <nil>, error",
			payment,
		)
	}
	t.Logf("message: %v", err)
}

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatalf(
			"GetPaymentMethod(Cash) = %v, %v, nil",
			payment,
			err,
		)
	}

	message := payment.Pay(10.30)
	want := "paid using cash"
	if !strings.Contains(message, want) {
		t.Errorf(
			"payment.Pay() = %q,\nshould contains %q, nil",
			message,
			want,
		)
	}

	t.Log("log:", message)
}

func TestCreatePaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatalf(
			"GetPaymentMethod(DebitCard) = %v, %v, nil",
			payment,
			err,
		)
	}

	message := payment.Pay(22.30)
	want := "paid using debit card"
	if !strings.Contains(message, want) {
		t.Errorf(
			"payment.Pay() = %q,\nwant match for %q, nil",
			message,
			want,
		)
	}

	t.Log("log:", message)
}
