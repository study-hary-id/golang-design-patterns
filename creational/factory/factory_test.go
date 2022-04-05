package factory

import (
	"strings"
	"testing"
)

// TestCreatePaymentMethodCash calls GetPaymentMethod with Cash option.
func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatalf("GetPaymentMethod() = %v, want match for Cash, nil", err)
	}

	message := payment.Pay(10.30)
	want := "paid using cash"
	if !strings.Contains(message, want) {
		t.Errorf(
			"payment.Pay() = '%v',\nwant match for '%v', nil",
			message,
			want,
		)
	}

	t.Log("log:", message)
}

// TestCreatePaymentMethodDebitCard calls GetPaymentMethod with DebitCard option.
func TestCreatePaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatalf("GetPaymentMethod() = %v, want match for Debit Card, nil", err)
	}

	message := payment.Pay(22.30)
	want := "paid using debit card"
	if !strings.Contains(message, want) {
		t.Errorf(
			"payment.Pay() = '%v',\nwant match for '%v', nil",
			message,
			want,
		)
	}

	t.Log("log:", message)
}

// TestGetPaymentMethodNonExistent calls GetPaymentMethod with wrong option.
func TestGetPaymentMethodNonExistent(t *testing.T) {
	payment, err := GetPaymentMethod(3)
	if err == nil {
		t.Errorf("GetPaymentMethod(3) = %v, want match for error, nil", payment)
	}
	t.Log("log:", err)
}
