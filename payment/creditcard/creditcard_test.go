package creditcard_test

import (
	"booklib/payment/creditcard"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	want := "123456789012"
	cc, err := creditcard.New(want)
	if err != nil {
		t.Fatal(err)
	}
	got := cc.Number()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestNewInvalid(t *testing.T) {
	t.Parallel()
	_, err := creditcard.New("")
	if err == nil {
		t.Fatal("want error for empty card number, but got nil")
	}
}
