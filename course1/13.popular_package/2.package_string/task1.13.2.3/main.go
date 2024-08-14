package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println(GenerateRandomString(10))
}

func GenerateRandomString(length int) string {
	rand.Seed(int64(time.Now().Nanosecond()))

	symbols := []rune("qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM")
	builder := strings.Builder{}
	builder.Grow(length * 2)
	for i := 0; i < length; i++ {
		num := rand.Intn(len(symbols) - 1)
		builder.WriteRune(symbols[num])
	}
	return builder.String()
}
