package service

import "dto"

func Topup(AccountNumber string, Amount int) dto.TopupResponse {
	return dto.TopupResponse{AccountName: "Ployploy",
		BalanceOld:     15000.00,
		CurrentBalance: 10000.00,
		Topup:          5000.00,
	}
}
