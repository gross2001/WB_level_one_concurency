package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	Способы остановки горутины:
 		все контексты (реализуем один?)
		канал (проверять закрыт ли)
*/

func quitWithChanel(quit chan struct{}) {
	for {
		select {
		case <-quit:
			fmt.Println("Goroutine quitWithChanel finishes...")
			return
		default:
			fmt.Println("Goroutine quitWithChanel works...")
			time.Sleep(1 * time.Second)
		}
	}
}

func quitWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine quitWithContext finishes...")
			return
		default:
			fmt.Println("Goroutine  quitWithContext works...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		quitWithContext(ctx)
	}()

	quit := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		quitWithChanel(quit)
	}()

	time.Sleep(1 * time.Second)
	close(quit)

	wg.Wait()
}
