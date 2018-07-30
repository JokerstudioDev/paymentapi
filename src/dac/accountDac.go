package dac

import "model"

type IAccountDac interface {
	UpdateBalance(accountNumber string, balance float64) bool
	ReadById(accountNumber string) model.Account
}
