package service

import (
	"dac"
	"dto"
	"testing"
)

func Test_Transfer_Input_AccountNumberStart_8976789543_AccountNumberDestination_8976909543_Amount_5000_Should_Be_Ployploy_BalanceOld_15000_CurrentBalance_20000_Amount_5000(t *testing.T) {
	expected := dto.TransferResponse{
		AccountName:    "Ployploy",
		BalanceOld:     15000.00,
		CurrentBalance: 10000.00,
		TransferAmount: 5000.00,
	}
	accountDac := dac.NewAccountDacFake()
	historyDac := dac.NewHistoryDacFake()
	paymentSvc := NewPaymentSvc(accountDac, historyDac)

	accualResult := paymentSvc.Transfer("8976789543", "8976909543", 5000.00)

	if expected != accualResult {
		t.Errorf("expected %v but got %v", expected, accualResult)
	}

}

func Test_Transfer_By_CalFee(t *testing.T) {
	expected := dto.TransferResponse{
		AccountName:    "Ployploy",
		BalanceOld:     15000.00,
		CurrentBalance: 8975.00,
		TransferAmount: 6025.00,
	}
	accountDac := dac.NewAccountDacFake()
	historyDac := dac.NewHistoryDacFake()
	paymentSvc := NewPaymentSvc(accountDac, historyDac)

	accualResult := paymentSvc.Transfer("8976789543", "8976909543", 6000.00)

	if expected != accualResult {
		t.Errorf("expected %v but got %v", expected, accualResult)
	}

}
