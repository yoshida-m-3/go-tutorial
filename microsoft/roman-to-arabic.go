package main

import "fmt"

// IV
// currentがnextより小さい値だったらcurrentを符号反転させる

func main() {
	romanList := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	var input string
	fmt.Scan(&input)

	// check
	for _, char := range input {
		if _, ok := romanList[char]; ok == false {
			fmt.Println("不正な文字が含まれています >> " + string(char))
			return
		}
	}

	// convert
	arabics := make([]int, len(input))

	for index, char := range input {
		arabics[index] = romanList[char]
	}

	// summary
	sum := 0

	for index, val := range arabics {
		if index < len(arabics)-1 && val < arabics[index+1] {
			val = val * -1
		}
		sum += val
	}

	fmt.Println(sum)
}
