package main

import (
	"errors"
	"sync"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

type SavingsAccount struct {
	balance float64
	mu      sync.Mutex
}

func (a *SavingsAccount) Deposit(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *SavingsAccount) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance-amount < 1000 {
		return errors.New("cannot withdraw: balance cannot be less than 1000")
	}
	a.balance -= amount
	return nil
}

func (a *SavingsAccount) Balance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

type CheckingAccount struct {
	balance float64
	mu      sync.Mutex
}

func (a *CheckingAccount) Deposit(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *CheckingAccount) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance < amount {
		return errors.New("cannot withdraw: insufficient funds")
	}
	a.balance -= amount
	return nil
}

func (a *CheckingAccount) Balance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

type Customer struct {
	Name    string
	Account Account
}

type Option func(*Customer)

func WithName(name string) Option {
	return func(c *Customer) {
		c.Name = name
	}
}

func WithAccount(account Account) Option {
	return func(c *Customer) {
		c.Account = account
	}
}

func NewCustomer(opts ...Option) *Customer {
	customer := &Customer{}
	for _, opt := range opts {
		opt(customer)
	}
	return customer
}
