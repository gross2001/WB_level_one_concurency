package ex02

import (
	"log"
	"sync"
)

func solutionByWaitGroup(src []int) {
	var wg sync.WaitGroup
	wg.Add(len(src))
	for _, numb := range src {
		go func() {
			defer wg.Done()
			log.Println(numb * numb)
		}()
	}
	wg.Wait()
}
