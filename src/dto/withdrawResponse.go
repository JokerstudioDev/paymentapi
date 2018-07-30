package dto

type WithdrawResponse struct {
	AccountName    string
	BalanceOld     float64
	CurrentBalance float64
	WithdrawAmount float64
}
