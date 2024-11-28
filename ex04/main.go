package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

var workersToReadNumb = 3

func workerToRead(wg *sync.WaitGroup, ch chan int, i int) {
	defer wg.Done()
	for {
		val, ok := <-ch
		if ok {
			fmt.Printf("read by %d value %d\n", i, val)
		} else {
			fmt.Printf("chanel is closed, nothing to read by %d\n", i)
			return
		}
	}
}

func workerToWrite(ctx context.Context, ch chan int) {
	var i int
writeChannel:
	for {
		select {
		case ch <- i:
			fmt.Printf("write value %d\n", i)
			i++
		case <-ctx.Done():
			fmt.Printf("finished main process\n")
			close(ch)
			break writeChannel
		}
	}
}

func run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ch := make(chan int, 5)
	var wg sync.WaitGroup
	for i := range workersToReadNumb {
		wg.Add(1)
		go workerToRead(&wg, ch, i)
	}

	workerToWrite(ctx, ch)

	wg.Wait()
}

func main() {
	run()
}
