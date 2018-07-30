package service

import (
	"dac"
	"dto"
	"testing"
)

func Test_Topup_Input_AccountNumberDestination_8976789543_Amount_5000_Should_Be_Ployploy_BalanceOld_15000_CurrentBalance_20000_Amount_5000(t *testing.T) {
	expected := dto.TopupResponse{
		AccountName:    "Ployploy",
		BalanceOld:     15000.00,
		CurrentBalance: 20000.00,
		Topup:          5000.00,
	}
	accountDac := dac.NewAccountDacFake()
	historyDac := dac.NewHistoryDacFake()
	paymentSvc := NewPaymentSvc(accountDac, historyDac)

	accualResult := paymentSvc.Topup("8976789543", 5000.00)

	if expected != accualResult {
		t.Errorf("expected %v but got %v", expected, accualResult)
	}

}
