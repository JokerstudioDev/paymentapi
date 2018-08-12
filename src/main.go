package main

import (
	"dac"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/topup", topupHandler)
	r.POST("/withdraw", withdrawHandler)
	r.POST("/transfer", transferHandler)
	r.Run()

}

func topupHandler(c *gin.Context) {
	type TopupReq struct {
		AccountNumber string  `json:"accountnumber"`
		Amount        float64 `json:"amount"`
	}
	var req TopupReq
	if err := c.ShouldBindJSON(&req); err == nil {
		accountDac := dac.NewAccountDac("root:123456@tcp(127.0.0.1:3306)/workshop")
		historyDac := dac.NewHistoryDac("root:123456@tcp(127.0.0.1:3306)/workshop")
		paymentSvc := service.NewPaymentSvc(accountDac, historyDac)
		results := paymentSvc.Topup(req.AccountNumber, req.Amount)

		c.JSON(200, results)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func withdrawHandler(c *gin.Context) {
	type WithdrawReq struct {
		AccountNumber string  `json:"accountnumber"`
		Amount        float64 `json:"amount"`
	}
	var req WithdrawReq
	if err := c.ShouldBindJSON(&req); err == nil {
		accountDac := dac.NewAccountDac("root:123456@tcp(127.0.0.1:3306)/workshop")
		historyDac := dac.NewHistoryDac("root:123456@tcp(127.0.0.1:3306)/workshop")
		paymentSvc := service.NewPaymentSvc(accountDac, historyDac)
		results := paymentSvc.Withdraw(req.AccountNumber, req.Amount)
		c.JSON(200, results)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func transferHandler(c *gin.Context) {
	type WithdrawReq struct {
		AccountNumberSource string  `json:"accountnumbersource"`
		AccountNumberDes    string  `json:"accountnumberdes"`
		Amount              float64 `json:"amount"`
	}
	var req WithdrawReq
	if err := c.ShouldBindJSON(&req); err == nil {
		accountDac := dac.NewAccountDac("root:123456@tcp(127.0.0.1:3306)/workshop")
		historyDac := dac.NewHistoryDac("root:123456@tcp(127.0.0.1:3306)/workshop")
		paymentSvc := service.NewPaymentSvc(accountDac, historyDac)
		results := paymentSvc.Transfer(req.AccountNumberSource, req.AccountNumberDes, req.Amount)
		c.JSON(200, results)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}
