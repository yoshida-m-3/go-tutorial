package main

import "fmt"

func fibonacci() func() int {
	prev := 0
	current := 1

	return func() int {
		prev, current = current, prev+current
		return prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
