package accounts

import "go_studies/oop/customers"

type CheckingAccount struct {
	Holder        customers.Holder
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (b *CheckingAccount) Withdraw(value float64) string {
	canWithdraw := value > 0 && value <= b.balance

	if canWithdraw {
		b.balance -= value
		return "withdrawal made successfully"
	}
	return "insufficient balance"
}

func (b *CheckingAccount) Deposit(value float64) string {
	if value > 0 {
		b.balance += value
		return "Successfully deposited"
	}
	return "error: value must be greater than zero"
}

func (b *CheckingAccount) Transfer(value float64, destinyAccount *CheckingAccount) bool {
	if value < b.balance && value > 0 {
		b.balance -= value
		destinyAccount.Deposit(value)
		return true
	} else {
		return false
	}
}

func (b *CheckingAccount) GetBalance() float64 {
	return b.balance
}
