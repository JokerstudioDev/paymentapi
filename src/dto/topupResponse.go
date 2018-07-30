package dto

type TopupResponse struct {
	AccountName    string
	BalanceOld     float64
	CurrentBalance float64
	Topup          float64
}
