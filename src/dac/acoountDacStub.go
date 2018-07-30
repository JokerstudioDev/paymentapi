package dac

import "model"

type AccountDacStub struct{}

func (*AccountDacStub) ReadById(accountNumber string) model.Account {
	return model.Account{ID: 1, AccountName: "Ployploy", AccountNumber: "8976789543", Balance: 15000.00}
}

func (*AccountDacStub) UpdateBalance(accountNumber string, balance float64) bool {
	return true
}
