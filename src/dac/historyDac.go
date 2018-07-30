package dac

import (
	"model"
)

type IHistoryDac interface {
	ReadByAccountNumber(accountNumber string) []model.History
	AddHistory(model.History) bool
}
