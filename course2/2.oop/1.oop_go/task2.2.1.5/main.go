package main

import "fmt"

type Mover interface {
	Move() string
	Speed() int
	MaxSpeed() int
	MinSpeed() int
}

type FastMover struct {
	BaseMover
}

func (f FastMover) Move() string {
	return fmt.Sprintf("Moving at speed: %d ", f.BaseMover.speed)
}

func (f FastMover) Speed() int {
	return f.BaseMover.speed
}

func (f FastMover) MaxSpeed() int {
	return f.BaseMover.speed + 20
}

func (f FastMover) MinSpeed() int {
	return f.BaseMover.speed - 90
}

type SlowMover struct {
	BaseMover
}

func (s SlowMover) Move() string {
	return fmt.Sprintf("Moving at speed: %d ", s.BaseMover.speed)
}

func (s SlowMover) Speed() int {
	return s.BaseMover.speed
}

func (s SlowMover) MaxSpeed() int {
	return s.BaseMover.speed + 110
}

func (s SlowMover) MinSpeed() int {
	return s.BaseMover.speed
}

type BaseMover struct {
	speed int
}

// Fast mover! Moving at speed: 100
// Maximum speed: 120
// Minimum speed: 10
// Slow mover...Moving at speed: 10
// Maximum speed: 120
// Minimum speed: 10

func main() {
	var movers []Mover

	fm := FastMover{BaseMover{100}}
	sm := SlowMover{BaseMover{10}}

	movers = append(movers, fm, sm)

	for _, mover := range movers {
		fmt.Println(mover.Move())
		fmt.Println("Maximum speed:", mover.MaxSpeed())
		fmt.Println("Minimum speed:", mover.MinSpeed())
	}
}
