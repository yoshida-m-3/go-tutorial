package main

import "fmt"

func add(x, y int) int { // 引数の型を省略できる
    return x +y
}

func main() {
    fmt.Println(add(42, 13))
}
