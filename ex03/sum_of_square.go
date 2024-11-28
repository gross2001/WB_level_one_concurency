package ex03

func square(numb int, result chan int) {
	result <- numb * numb

}

func sumOfSquare(src []int) (sum int) {
	ch := make(chan int, len(src))
	for _, numb := range src {
		go square(numb, ch)
	}

	for range src {
		val := <-ch
		sum += val
	}
	return
}
