package ex18

import (
	"sync"
)

type counterStruct struct {
	counter int
	mu      sync.Mutex
}

func increment(wg *sync.WaitGroup, counter *counterStruct) {
	defer counter.mu.Unlock()
	defer wg.Done()
	counter.mu.Lock()
	counter.counter++
}

func run(howManyIncrement int) int {
	if howManyIncrement <= 0 {
		return 0
	}
	var counter counterStruct
	var wg sync.WaitGroup

	for range howManyIncrement {
		wg.Add(1)
		go increment(&wg, &counter)
	}

	wg.Wait()
	return counter.counter
}
