package service

import (
	"dac"
	"dto"
	"fmt"
	"model"
	"time"
)

type PaymentSvc struct{}

var accountDac dac.IAccountDac
var historyDac dac.IHistoryDac

func NewPaymentSvc(accountDacObj dac.IAccountDac, historyDacObj dac.IHistoryDac) *PaymentSvc {
	accountDac = accountDacObj
	historyDac = historyDacObj
	return &PaymentSvc{}
}

func (*PaymentSvc) Topup(accountNumber string, amount float64) dto.TopupResponse {
	account := accountDac.ReadById(accountNumber)
	balance := account.Balance + amount

	transactionResult := accountDac.UpdateBalance(accountNumber, balance)

	if transactionResult == false {
		historyDac.AddHistory(model.History{
			AccountNumber: account.AccountNumber,
			Date:          time.Now(),
			Description:   fmt.Sprintf("Topup %f Bath fail", amount),
		})
		return dto.TopupResponse{}
	}

	historyDac.AddHistory(model.History{
		AccountNumber: account.AccountNumber,
		Date:          time.Now(),
		Description:   fmt.Sprintf("Topup %f Bath success", amount),
	})

	return dto.TopupResponse{AccountName: account.AccountName,
		BalanceOld:     account.Balance,
		CurrentBalance: balance,
		Topup:          amount,
	}
}
