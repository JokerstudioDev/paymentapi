package main

import (
	"dac"
	"fmt"
	"service"
)

func main() {
	accountDac := dac.NewAccountDac("root:123456@tcp(127.0.0.1:3306)/workshop")
	historyDac := dac.NewHistoryDacFake()
	paymentSvc := service.NewPaymentSvc(accountDac, historyDac)
	results := paymentSvc.Topup("ploy01", 5000.00)
	fmt.Println(results)
}
