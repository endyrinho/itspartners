package wallet_project_test

import (
	"fmt"
	"testing"
	"wallet_project"
)

func TestWallet_Deposit(t *testing.T) {
	wallet := wallet_project.NewWallet()
	wallet.Deposit(10)

	assertBalance(t, wallet, 10)
}

func TestWallet_WithdrawSufficientFunds(t *testing.T) {
	wallet := wallet_project.NewWallet()
	wallet.Deposit(20)
	err := wallet.Withdraw(10)

	assertNoError(t, err)
	assertBalance(t, wallet, 10)
}

func TestWallet_WithdrawInsufficientFunds(t *testing.T) {
	wallet := wallet_project.NewWallet()
	wallet.Deposit(20)
	err := wallet.Withdraw(30)

	assertError(t, err, "insufficient funds")
	assertBalance(t, wallet, 20)
}

func assertBalance(t *testing.T, wallet *wallet_project.Wallet, want float64) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Balance was incorrect, got: %f, want: %f", got, want)
	} else {
		fmt.Printf("Balance test passed: got: %f, want: %f\n", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else {
		fmt.Println("No error occurred.")
	}
}

func assertError(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("Expected an error but didn't get one")
	}

	if got.Error() != want {
		t.Errorf("Error message was incorrect, got: %q, want: %q", got, want)
	} else {
		fmt.Printf("Error test passed: got: %q, want: %q\n", got, want)
	}
}
