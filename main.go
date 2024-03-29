package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/mfdrfauzi/fwg17-go-introduction-3/app"
)

func main() {
	//Sumerizing with Channel
	fmt.Println("---Sumerizing with Channel---")
	a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	chn := make(chan int)

	go app.Sum(a[:len(a)/2], chn) //Result 1
	go app.Sum(a[len(a)/2:], chn) //Result 2

	result1 := <-chn
	result2 := <-chn
	total := result1 + result2
	fmt.Println("Result 1:", result1)
	fmt.Println("Result 2:", result2)
	fmt.Println("Total Sum:", total)
	time.Sleep(1 * time.Second)

	//Deret with Goroutine
	deret := app.DeretBilangan{Limit: 40}
	fmt.Println("\n---Deret with Go-Routine---")
	go deret.Prima()
	go deret.Ganjil()
	go deret.Genap()
	go deret.Fibonacci()
	time.Sleep(2 * time.Second)

	//Functions with Goroutine and WaitGroup
	fmt.Println("\n---Data Trade with GOroutine and Waitgroup---")
	var wg sync.WaitGroup

	fibonacciResultCh := make(chan int, 10)
	oddCh := make(chan int, 10)
	evenCh := make(chan int, 10)

	wg.Add(1)
	go app.Fibonacci(30, fibonacciResultCh, &wg)

	wg.Add(1)
	go app.OddEven(fibonacciResultCh, oddCh, evenCh, &wg)

	wg.Wait()

	fmt.Print("Bilangan Fibonacci: ")
	for fib := range fibonacciResultCh {
		fmt.Print(fib, " ")
	}

	fmt.Print("\nBilangan Ganjil: ")
	for odd := range oddCh {
		fmt.Print(odd, " ")
	}

	fmt.Print("\nBilangan Genap:")
	for even := range evenCh {
		fmt.Print(even, " ")
	}
	time.Sleep(3 * time.Second)

	//Program Aplikasi Pembagian Jadwal Piket
	fmt.Println("\n\n---Pembagian Jadwal Piket---")
	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}
	persons := []string{"Dzulfikar", "Fauzi", "Gabriel", "David", "Eva"}

	scheduleCh := make(chan app.Schedule, 5)

	wg.Add(2)

	go app.GenerateSchedule(days, persons, scheduleCh, &wg)
	go app.PrintSchedule(scheduleCh, &wg)

	wg.Wait()
	fmt.Println("\nProgram selesai.")
}
