package main

import "fmt"

func split(sum int) (x, y int) { // 戻り値の変数名を定義可能
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}
