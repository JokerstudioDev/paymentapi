package service

type TopupResponse struct {
	AccountName    string
	BalanceOld     float64
	CurrentBalance float64
	Topup          float64
}

func Topup(AccountNumber string, Amount int) TopupResponse {
	return TopupResponse{AccountName: "Ployploy",
		BalanceOld:     15000.00,
		CurrentBalance: 10000.00,
		Topup:          5000.00,
	}
}
