package main

import (
	"math/rand"
	"strings"
	"time"
)

func generateActivationKey() string {
	rand.Seed(int64(time.Now().Nanosecond()))

	parts := make([]string, 4)
	part := make([]string, 4)

	for i := range parts {
		for j := range part {

			if (rand.Int() & 1) == 0 {
				part[j] = string(byte(rand.Intn(10) + 48))
				continue
			}
			part[j] = string(byte(rand.Intn(26) + 65))
		}
		parts[i] = strings.Join(part, "")
	}
	return strings.Join(parts, "-")
}
