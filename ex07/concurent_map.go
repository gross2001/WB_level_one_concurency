package main

import (
	"sync"
)

type mapConcurent struct {
	sync.RWMutex
	m map[int]int
}

func writer(src *mapConcurent, key, value int) {
	src.Lock()
	defer src.Unlock()
	src.m[key] = value
}
