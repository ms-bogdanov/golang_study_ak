package main

import (
	"testing"
)

func TestSavingsAccount_Deposit(t *testing.T) {
	account := &SavingsAccount{}
	account.Deposit(2000)
	if account.Balance() != 2000 {
		t.Errorf("Expected balance 2000, got %v", account.Balance())
	}
}

func TestSavingsAccount_Withdraw(t *testing.T) {
	account := &SavingsAccount{}
	account.Deposit(2000)
	err := account.Withdraw(500)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if account.Balance() != 1500 {
		t.Errorf("Expected balance 1500, got %v", account.Balance())
	}

	err = account.Withdraw(600)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if account.Balance() != 1500 {
		t.Errorf("Expected balance 1500, got %v", account.Balance())
	}
}

func TestSavingsAccount_Balance(t *testing.T) {
	account := &SavingsAccount{}
	account.Deposit(2000)
	if account.Balance() != 2000 {
		t.Errorf("Expected balance 2000, got %v", account.Balance())
	}
}

func TestCheckingAccount_Deposit(t *testing.T) {
	account := &CheckingAccount{}
	account.Deposit(2000)
	if account.Balance() != 2000 {
		t.Errorf("Expected balance 2000, got %v", account.Balance())
	}
}

func TestCheckingAccount_Withdraw(t *testing.T) {
	account := &CheckingAccount{}
	account.Deposit(2000)
	err := account.Withdraw(500)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if account.Balance() != 1500 {
		t.Errorf("Expected balance 1500, got %v", account.Balance())
	}

	err = account.Withdraw(2000)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if account.Balance() != 1500 {
		t.Errorf("Expected balance 1500, got %v", account.Balance())
	}
}

func TestCheckingAccount_Balance(t *testing.T) {
	account := &CheckingAccount{}
	account.Deposit(2000)
	if account.Balance() != 2000 {
		t.Errorf("Expected balance 2000, got %v", account.Balance())
	}
}

func TestNewCustomer(t *testing.T) {
	account := &SavingsAccount{}
	account.Deposit(2000)

	customer := NewCustomer(
		WithName("John Doe"),
		WithAccount(account),
	)

	if customer.Name != "John Doe" {
		t.Errorf("Expected name John Doe, got %v", customer.Name)
	}

	if customer.Account.Balance() != 2000 {
		t.Errorf("Expected balance 2000, got %v", customer.Account.Balance())
	}
}

func TestWithName(t *testing.T) {
	customer := &Customer{}
	WithName("John Doe")(customer)

	if customer.Name != "John Doe" {
		t.Errorf("Expected name John Doe, got %v", customer.Name)
	}
}

func TestWithAccount(t *testing.T) {
	account := &SavingsAccount{}
	account.Deposit(2000)

	customer := &Customer{}
	WithAccount(account)(customer)

	if customer.Account.Balance() != 2000 {
		t.Errorf("Expected balance 2000, got %v", customer.Account.Balance())
	}
}
