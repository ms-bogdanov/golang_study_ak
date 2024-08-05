package main

import (
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID       int
	Complete bool
}

var orders []*Order
var completeOrders map[int]bool
var wg sync.WaitGroup
var processTimes chan time.Duration
var sinceProgramStarted time.Duration
var count int
var limitCount int

func main() {
	count = 30
	limitCount = 5
	processTimes = make(chan time.Duration, count)
	orders = GenerateOrders(count)
	completeOrders = GenerateCompleteOrders(count)
	programStart := time.Now()
	LimitSpawnOrderProcessing(limitCount)

	wg.Wait()
	sinceProgramStarted = time.Since(programStart)

	go func() {
		time.Sleep(1 * time.Second)
		close(processTimes)
	}()

	checkTimeDifference(limitCount)
}

func checkTimeDifference(limitCount int) {
	var averageTime time.Duration
	var orderProcessTotalTime time.Duration
	var orderProcessedCount int

	for v := range processTimes {
		orderProcessedCount++
		orderProcessTotalTime += v
	}

	if orderProcessedCount != count {
		panic("orderProcessedCount != count")
	}

	averageTime = orderProcessTotalTime / time.Duration(orderProcessedCount)
	println("orderProcessTotalTime", orderProcessTotalTime/time.Second)
	println("averageTime", averageTime/time.Second)
	println("sinceProgramStarted", sinceProgramStarted/time.Second)
	println("sinceProgramStarted average", sinceProgramStarted/(time.Duration(orderProcessedCount)*time.Second))
	println("orderProcessTotalTime - sinceProgramStarted", (orderProcessTotalTime-sinceProgramStarted)/time.Second)

	if (orderProcessTotalTime/time.Duration(limitCount)-sinceProgramStarted)/time.Second > 0 {
		panic("(orderProcessTotalTime-sinceProgramStarted)/time.Second > 0")
	}
}

func OrderProcessing(order *Order, limit chan struct{}, t time.Time) {
	if completeOrders[order.ID] {
		processTimes <- time.Since(t)
		<-limit
		wg.Done()
		return
	}

	time.Sleep(1 * time.Second)
	processTimes <- time.Since(t)
	<-limit
	wg.Done()
}

func LimitSpawnOrderProcessing(limitCount int) {
	limit := make(chan struct{}, limitCount)
	var t time.Time

	for _, i := range orders {
		limit <- struct{}{}
		wg.Add(1)
		t = time.Now()
		go OrderProcessing(i, limit, t)
	}
}

func GenerateOrders(count int) []*Order {
	result := []*Order{}
	for i := 0; i < count; i++ {
		newOrder := &Order{ID: i, Complete: false}
		result = append(result, newOrder)
	}

	return result
}

func GenerateCompleteOrders(maxOrderID int) map[int]bool {
	result := make(map[int]bool)
	for _, i := range orders {
		randomNumber := rand.Float64()

		if randomNumber > 0.5 {
			result[i.ID] = true
		}
	}

	return result
}
