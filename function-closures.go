package main

import "fmt"

func adder() func(int) int {
	sum := 0 // pos, negから共通で参照されず別々
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
