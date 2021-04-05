package accounts

import "go_studies/oop/customers"

type SavingAccount struct {
	Holder                                 customers.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (s *SavingAccount) Withdraw(value float64) string {
	canWithdraw := value > 0 && value <= s.balance

	if canWithdraw {
		s.balance -= value
		return "withdrawal made successfully"
	}
	return "insufficient balance"
}

func (s *SavingAccount) Deposit(value float64) string {
	if value > 0 {
		s.balance += value
		return "Successfully deposited"
	}
	return "error: value must be greater than zero"
}

func (s *SavingAccount) Transfer(value float64, destinyAccount *SavingAccount) bool {
	if value < s.balance && value > 0 {
		s.balance -= value
		destinyAccount.Deposit(value)
		return true
	} else {
		return false
	}
}

func (s *SavingAccount) GetBalance() float64 {
	return s.balance
}
