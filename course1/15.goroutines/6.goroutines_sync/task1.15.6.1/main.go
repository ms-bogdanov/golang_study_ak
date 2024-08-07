package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	count := 1000
	goroutines := make([]func() string, count)
	for i := 0; i < count; i++ {
		j := i
		goroutines[i] = func() string {
			return fmt.Sprintf("goroutine %d done", j)
		}
	}
	fmt.Println(waitGroupExample(goroutines...))
}

func waitGroupExample(goroutines ...func() string) string {
	wg := sync.WaitGroup{}
	res := make([]string, len(goroutines))

	for id, goroutine := range goroutines {
		wg.Add(1)
		go func(id int, goroutine func() string) {
			defer wg.Done()
			res[id] = goroutine()
		}(id, goroutine)
	}

	wg.Wait()

	result := strings.Join(res, "\n")
	return result
}
