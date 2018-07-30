package transfer

import (
	"testing"
)

func Test_TransferMoney_Input_Chonnikan_6041072620_500_Should_be_9000(t *testing.T) {
	expected := 9000
	accountName := "chonnikan"
	accountNaumber := "6041072620"
	amountTransfer := 500

	accual := TransferMoney(accountName, accountNaumber, amountTransfer)

	if accual != expected {
		t.Errorf("expected %d but got %d", expected, accual)
	}

}
