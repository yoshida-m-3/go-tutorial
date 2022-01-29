package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	result := fibonacci(number)

	fmt.Println(result)
}

func fibonacci(number int) []int {
	if number < 2 {
		return nil
	}

	sequences := []int{1, 1}

	for number > len(sequences) {
		sum := 0
		for _, v := range sequences[len(sequences)-2:] {
			sum += v
		}
		sequences = append(sequences, sum)
	}

	return sequences
}
