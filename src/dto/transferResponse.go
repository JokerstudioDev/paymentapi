package dto

type TransferResponse struct {
	AccountName    string
	BalanceOld     float64
	CurrentBalance float64
	TransferAmount float64
}
