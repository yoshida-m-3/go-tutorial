package main

import "fmt"

var i, j int = 1, 2 // 宣言と同時に初期化可能

func main() {
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)
}
