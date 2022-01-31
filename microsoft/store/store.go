package store

import (
	"errors"
	"fmt"
)

type Account struct {
	FirstName string
	LastName  string
}

func (a *Account) ChangeName(firstName string, lastName string) {
	a.FirstName = firstName
	a.LastName = lastName
}

type Employee struct {
	Account
	Credits float64
}

func (e *Employee) String() string {
	return fmt.Sprintf("Name: %s %s\nCredits: %.2f\n", e.FirstName, e.LastName, e.Credits)
}

func CreateEmployee(firstName, lastName string, credits float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credits}, nil
}

func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		e.Credits += amount
		return e.Credits, nil
	}

	return 0.0, errors.New("金額が正しくありません。")
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.Credits {
			e.Credits -= amount
			return e.Credits, nil
		} else {
			return 0.0, errors.New("残高以上に差し引けません。")
		}
	}

	return 0.0, errors.New("金額が正しくありません。")
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}

func main() {
	account := Account{"choco", "ball"}

	account.ChangeName("big", "ball")
	fmt.Println(account)
}
