package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b bool
	var a int
	var c int8
	var d int16
	var e int32
	var f int64
	var g uint
	var h uint8

	fmt.Println("bool size: ", sizeOfBool(b))
	fmt.Println("int size: ", sizeOfInt(a))
	fmt.Println("int8 size: ", sizeOfInt8(c))
	fmt.Println("int16 size: ", sizeOfInt16(d))
	fmt.Println("int32 size: ", sizeOfInt32(e))
	fmt.Println("int64 size: ", sizeOfInt64(f))
	fmt.Println("uint size: ", sizeOfUint(g))
	fmt.Println("uint8 size: ", sizeOfUint8(h))
}

func sizeOfBool(b bool) int {
	return int(unsafe.Sizeof(b))
}

func sizeOfInt(n int) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt8(n int8) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt16(n int16) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt32(n int32) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfInt64(n int64) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfUint(n uint) int {
	return int(unsafe.Sizeof(n))
}

func sizeOfUint8(n uint8) int {
	return int(unsafe.Sizeof(n))
}
