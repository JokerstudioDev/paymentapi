package dac

import "model"

type AccountDacFake struct{}

func NewAccountDacFake() *AccountDacFake {
	return &AccountDacFake{}
}

var accounts = []model.Account{
	model.Account{ID: 1, AccountName: "Ployploy", AccountNumber: "8976789543", Balance: 15000.00},
	model.Account{ID: 1, AccountName: "Joker", AccountNumber: "8976909543", Balance: 45900.00},
}

func (*AccountDacFake) ReadById(accountNumber string) model.Account {
	account := model.Account{}
	for _, u := range accounts {
		if u.AccountNumber == accountNumber {
			account = u
		}
	}
	return account
}

func (*AccountDacFake) UpdateBalance(accountNumber string, balance float64) bool {
	return true
}
