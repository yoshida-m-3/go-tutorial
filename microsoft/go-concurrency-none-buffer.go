package main

import (
	"fmt"
	"time"
)

var quit = make(chan bool)

func fib(ch chan int) {
	x, y := 1, 1

	for {
		select { // 2つの条件が実行されるまでスリープする
		case ch <- x:
			x, y = y, x+y
		case <-quit: // quitチャネルにセットされるまで待機するため止まる
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}

func main() {
	start := time.Now()
	command := ""

	data := make(chan int)

	go fib(data)

	for {
		num := <-data
		fmt.Println(num)

		fmt.Scanf("%s", &command) // ユーザーの入力受付のため一時停止する

		if command == "quit" {
			quit <- true
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
