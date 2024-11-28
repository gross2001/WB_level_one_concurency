package main

import (
	"context"
	"fmt"
	"time"
)

func run() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch := make(chan int, 1)
	i := 0

	for {
		select {
		case ch <- i:
			fmt.Println("Write new value", i)
			time.Sleep(100 * time.Millisecond)
			i++
		case val := <-ch:
			fmt.Println("Read new value", val)
			time.Sleep(100 * time.Millisecond)
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		}
	}
}

func main() {
	run()
}
