package service

import "testing"

func Test_Topup_Input_AccountNumberDestination_8976789543_Amount_5000_Should_Be_Ployploy_BalanceOld_15000_CurrentBalance_10000_Amount_5000(t *testing.T) {
	expected := TopupResponse{
		AccountName:    "Ployploy",
		BalanceOld:     15000.00,
		CurrentBalance: 10000.00,
		Topup:          5000.00,
	}

	accualResult := Topup("8976789543", 5000)

	if expected != accualResult {
		t.Errorf("expected %v but got %v", expected, accualResult)
	}

}
