package service

import (
	"dac"
	"dto"
	"testing"
)

func Test_Withdraw_Input_AccountNumber_8976909543_Amount_5000_Should_Be_Joker_BalanceOld_45900_CurrentBalance_40900_Amount_5000(t *testing.T) {
	expected := dto.WithdrawResponse{
		AccountName:    "Joker",
		BalanceOld:     45900.00,
		CurrentBalance: 40900.00,
		WithdrawAmount: 5000.00,
	}
	accountDac := dac.NewAccountDacFake()
	historyDac := dac.NewHistoryDacFake()
	paymentSvc := NewPaymentSvc(accountDac, historyDac)
	accualResult := paymentSvc.Withdraw("8976909543", 5000.00)

	if expected != accualResult {
		t.Errorf("expected %v but got %v", expected, accualResult)
	}
}
