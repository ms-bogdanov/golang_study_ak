package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type OrderOption func(*Order)

func WithCustomerID(CustomerID string) OrderOption {
	return func(o *Order) {
		o.CustomerID = CustomerID
	}
}

func WithItems(Items []string) OrderOption {
	return func(o *Order) {
		o.Items = Items
	}
}

func WithOrderDate(OrderDate time.Time) OrderOption {
	return func(o *Order) {
		o.OrderDate = OrderDate
	}
}

func NewOrder(ID int, options ...OrderOption) *Order {
	ord := &Order{
		ID: ID,
	}

	for _, option := range options {
		option(ord)
	}
	return ord
}

func main() {
	order := NewOrder(1,
		WithCustomerID("123"),
		WithItems([]string{"item1", "item2"}),
		WithOrderDate(time.Now()))

	fmt.Printf("Order: %+v\n", order)
}
