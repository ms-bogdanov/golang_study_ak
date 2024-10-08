package main

import (
	"errors"
	"fmt"
)

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}

type DineInOrder struct {
	orderDetails map[string]int
}

func (d *DineInOrder) AddItem(item string, quantity int) error {
	if quantity <= 0 {
		return errors.New("Error")
	}

	d.orderDetails[item] += quantity

	return nil
}

func (d *DineInOrder) RemoveItem(item string) error {
	delete(d.orderDetails, item)

	return nil
}

func (d *DineInOrder) GetOrderDetails() map[string]int {
	return d.orderDetails
}

type TakeAwayOrder struct {
	orderDetails map[string]int
}

func (t *TakeAwayOrder) AddItem(item string, quantity int) error {
	if quantity <= 0 {
		return errors.New("Error")
	}

	t.orderDetails[item] += quantity

	return nil
}

func (t *TakeAwayOrder) RemoveItem(item string) error {
	delete(t.orderDetails, item)

	return nil
}

func (t *TakeAwayOrder) GetOrderDetails() map[string]int {
	return t.orderDetails
}
func ManageOrder(o Order) {
	o.AddItem("Pizza", 2)
	o.AddItem("Burger", 1)
	o.RemoveItem("Pizza")
	fmt.Println(o.GetOrderDetails())
}

func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}

	fmt.Println("DineIn Order:")
	ManageOrder(dineIn)

	fmt.Println("TakeAway Order:")
	ManageOrder(takeAway)
}
