package dac

import (
	"model"
	"time"
)

type HistoryDacFake struct{}

func NewHistoryDacFake() *HistoryDacFake {
	return &HistoryDacFake{}
}

var histories = []model.History{
	model.History{ID: 1, AccountNumber: "8976789543", Date: time.Now(), Description: "Topup 5000.00 Bath success"},
	model.History{ID: 2, AccountNumber: "8976789543", Date: time.Now(), Description: "Topup 10000.00 Bath fail"},
}

func (*HistoryDacFake) ReadByAccountNumber(accountNumber string) []model.History {
	history := []model.History{}
	for _, h := range histories {
		if h.AccountNumber == accountNumber {
			history = append(history, h)
		}
	}
	return history
}

func (*HistoryDacFake) AddHistory(model.History) bool {
	return true
}
