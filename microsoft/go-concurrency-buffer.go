package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(number float64, ch chan string) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	// delay
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- fmt.Sprintf("Fib(%v): %v\n", number, x)
}

func main() {
	start := time.Now()
	size := 15
	ch := make(chan string, size)

	for i := 1; i < size; i++ {
		go fib(float64(i), ch)
	}

	for i := 1; i < size; i++ {
		fmt.Printf("Fib(%v): %v\n", i, <-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
