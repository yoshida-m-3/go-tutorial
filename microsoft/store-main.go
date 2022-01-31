package main

import (
	"fmt"
	"go-tutorial/microsoft/store"
)

func main() {
	kyoro, _ := store.CreateEmployee("choco", "ball", 100)
	fmt.Println(kyoro.CheckCredits())

	credits, err := kyoro.AddCredits(20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}

	_, err = kyoro.RemoveCredits(1000)

	if err != nil {
		fmt.Println("残高以上は引き出せません。", err)
	}

	kyoro.ChangeName("white choco", "ball")
	fmt.Println(kyoro)
}
