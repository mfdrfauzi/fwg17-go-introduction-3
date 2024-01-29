package app

import (
	"sync"
)

func Fibonacci(limit int, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	a, b := 0, 1
	for a <= limit {
		resultCh <- a
		a, b = b, a+b
	}

	close(resultCh)
}

func OddEven(numbersCh <-chan int, oddCh chan<- int, evenCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range numbersCh {
		if num%2 == 0 {
			evenCh <- num
		} else {
			oddCh <- num
		}
	}

	close(oddCh)
	close(evenCh)
}
