package main

import (
	"fmt"
)

func writer(src [10]int, ch chan int) {
	for i := range len(src) {
		ch <- src[i]
	}
	close(ch)
}

func computer(in, out chan int) {
	for {
		val, ok := <-in
		if ok {
			square := val * val
			out <- square
		} else {
			close(out)
			break
		}
	}
}

func printer(in chan int, quit chan struct{}) {
	for {
		val, ok := <-in
		if ok {
			fmt.Println(val)
		} else {
			break
		}
	}
	quit <- struct{}{}
}

func run() {
	src := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	chForNumb := make(chan int, 1)
	chForSqrt := make(chan int, 1)
	quit := make(chan struct{}, 1)

	go writer(src, chForNumb)
	go computer(chForNumb, chForSqrt)
	go printer(chForSqrt, quit)

	<-quit
}

func main() {
	run()
}
