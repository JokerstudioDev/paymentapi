package dac

import "model"

type AccountDacFake struct {
	UpdateBalanceCount int
	ReadByIDCount      int
}

func NewAccountDacFake() *AccountDacFake {
	return &AccountDacFake{UpdateBalanceCount: 0, ReadByIDCount: 0}
}

var accounts = []model.Account{
	model.Account{ID: 1, AccountName: "Ployploy", AccountNumber: "8976789543", Balance: 15000.00},
	model.Account{ID: 1, AccountName: "Joker", AccountNumber: "8976909543", Balance: 45900.00},
}

func (accountDac *AccountDacFake) ReadById(accountNumber string) model.Account {
	accountDac.ReadByIDCount += 1
	account := model.Account{}
	for _, u := range accounts {
		if u.AccountNumber == accountNumber {
			account = u
		}
	}
	return account
}

func (accountDac *AccountDacFake) UpdateBalance(accountNumber string, balance float64) bool {
	accountDac.UpdateBalanceCount += 1
	return true
}
