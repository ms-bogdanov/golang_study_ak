package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"runtime"
	"time"
)

func monitorGorutines(prevGorutines int) {
	currentGorutines := prevGorutines

	for {
		time.Sleep(300 * time.Millisecond)
		goroutinesNow := runtime.NumGoroutine()
		goroutinesDiff := float64(goroutinesNow) / float64(currentGorutines)

		if goroutinesDiff < 0.8 {
			fmt.Printf("Предупреждение: Количество горутин уменьшилось более чем на 20%%:\nТекущее количество горутин: %d", goroutinesNow)
		}

		if goroutinesDiff > 1.2 {
			fmt.Printf("Предупреждение: Количество горутин увеличилось более чем на 20%%:\nТекущее количество горутин: %d", goroutinesNow)
		}

		fmt.Printf("Текущее количество горутин: %d\n", goroutinesNow)
		currentGorutines = goroutinesNow
	}

}

func main() {
	var g, _ = errgroup.WithContext(context.Background())

	go func() {
		monitorGorutines(runtime.NumGoroutine())
	}()

	for i := 0; i < 64; i++ {
		g.Go(func() error {
			time.Sleep(5 * time.Second)
			return nil
		})
		time.Sleep(80 * time.Second)

		if err := g.Wait(); err != nil {
			fmt.Println(err)
		}
	}
}
