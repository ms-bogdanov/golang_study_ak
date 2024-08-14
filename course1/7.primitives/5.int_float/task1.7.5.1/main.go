package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println(binaryStringToFloat("00111110001000000000000000000000"))
}

func binaryStringToFloat(binary string) float32 {
	var number uint64
	number, _ = strconv.ParseUint(binary, 2, 64)
	floatNumber := *(*float32)(unsafe.Pointer(&number))
	return floatNumber
}
