package main

import (
	"time"
)

func customSleep(x time.Duration) {
	<-time.After(x)
}

func main() {
	customSleep(10 * time.Second)
}
