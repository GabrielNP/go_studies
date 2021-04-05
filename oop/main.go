package main

import (
	"fmt"
	"go_studies/oop/accounts"
)

func payBillet(account verifyAccount, value float64) {
	account.Withdraw((value))
}

type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {
	accountA := accounts.CheckingAccount{}
	accountA.Deposit(500)
	payBillet(&accountA, 440)

	fmt.Println(accountA.GetBalance())

	accountB := accounts.SavingAccount{}
	accountB.Deposit(500)
	payBillet(&accountB, 1000)

	fmt.Println(accountB.GetBalance())
}
