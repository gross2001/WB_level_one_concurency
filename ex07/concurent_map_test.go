package main

import (
	"reflect"
	"sync"
	"testing"
)

func Test_writer(t *testing.T) {
	tests := []struct {
		name       string
		iterations int
		result     map[int]int
	}{
		{
			name:       "first",
			iterations: 5,
			result:     map[int]int{0: 0, 1: 1, 2: 4, 3: 9, 4: 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := mapConcurent{
				m: make(map[int]int),
			}
			var wg sync.WaitGroup

			for i := range tt.iterations {
				wg.Add(1)
				go func() {
					defer wg.Done()
					writer(&src, i, i*i)
				}()
			}
			wg.Wait()

			eq := reflect.DeepEqual(src.m, tt.result)
			if !eq {
				t.Fatalf("Expect map %v, got %v", tt.result, src.m)
			}
		})
	}
}
