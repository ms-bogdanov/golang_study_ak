package main

import (
	"container/list"
	"fmt"
)

type Car struct {
	LicensePlate string
}

type ParkingLot struct {
	space *list.List
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{space: list.New()}
}

func (p *ParkingLot) Park(c Car) {
	fmt.Println("Car number ", c.LicensePlate, " is parked.")
	p.space.PushBack(c)
}

func (p *ParkingLot) Leave() {
	firstCar := p.space.Front()

	if firstCar == nil {
		fmt.Println("Car parking is empty")
		return
	}

	fmt.Println("Car number: ", firstCar.Value.(Car).LicensePlate, " has left parking")
	p.space.Remove(firstCar)
}

func main() {
	parkingLot := NewParkingLot()
	parkingLot.Park(Car{LicensePlate: "ABC-123"})
	parkingLot.Park(Car{LicensePlate: "XYZ-789"})
	parkingLot.Leave()
	parkingLot.Leave()
	parkingLot.Leave()
}
