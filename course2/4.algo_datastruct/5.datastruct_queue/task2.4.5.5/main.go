package main

import "fmt"

type RingBuffer struct {
	buffer        []int
	lastElemIndex int
	len           int
	cap           int
	addFlag       int
}

func NewRingBuffer(cap int) *RingBuffer {
	bufferSlice := make([]int, cap)

	return &RingBuffer{
		buffer:        bufferSlice,
		lastElemIndex: cap - 1,
		len:           0,
		addFlag:       0,
		cap:           cap,
	}
}

type CircuitRinger interface {
	Add(val int)
	Get() (int, bool)
}

func (buffer *RingBuffer) Add(val int) {
	buffer.buffer[buffer.addFlag] = val
	if buffer.len < buffer.cap {
		buffer.len++
	}

	if buffer.lastElemIndex == buffer.addFlag {
		buffer.addFlag = 0
	} else {
		buffer.addFlag++
	}
}

func (buffer *RingBuffer) Get() (int, bool) {
	if buffer.len == 0 {
		return 0, false
	}

	getFlag := buffer.addFlag - buffer.len

	if getFlag < 0 {
		getFlag = buffer.cap + getFlag
	}

	currentElem := buffer.buffer[getFlag]
	buffer.len--
	return currentElem, true
}

func main() {
	rb := NewRingBuffer(3)
	rb.Add(1)
	rb.Add(2)
	rb.Add(3)
	rb.Add(4)

	for val, ok := rb.Get(); ok; val, ok = rb.Get() {
		fmt.Println(val)
	}

	if _, ok := rb.Get(); !ok {
		fmt.Println("Buffer is empty")
	}
}
