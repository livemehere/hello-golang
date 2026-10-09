package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%d canceled\n", id)
			return

		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("%d working...\n", id)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := range 3 {
		go worker(ctx, i)
	}

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(time.Second)
}
