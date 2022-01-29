package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case isFizz(i) && isBuzz(i):
			fmt.Println("FizzBuzz")
		case isFizz(i):
			fmt.Println("Fizz")
		case isBuzz(i):
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func isFizz(num int) bool {
	return num%3 == 0
}

func isBuzz(num int) bool {
	return num%5 == 0
}
