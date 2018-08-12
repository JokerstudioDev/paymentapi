package dac

import (
	"database/sql"
	"fmt"
	"model"

	_ "github.com/go-sql-driver/mysql"
)

type AccountDac struct{}

var connString string

func NewAccountDac(conn string) *AccountDac {
	connString = conn
	return &AccountDac{}
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Println("connection fails")
		panic(err)
	}
	fmt.Println("connect success")
	return db
}

func (*AccountDac) UpdateBalance(accountNumber string, balance float64) bool {
	db := openDB()
	defer db.Close()
	statement, _ := db.Prepare("UPDATE `account` SET balance= ? WHERE account_number=?")
	defer statement.Close()
	_, err := statement.Exec(balance, accountNumber)
	if err != nil {
		return false
	}
	return true
}

func (*AccountDac) ReadById(accountNumber string) model.Account {
	db := openDB()
	defer db.Close()
	results, _ := db.Query("SELECT * FROM account WHERE account_number=?", accountNumber)
	var account model.Account
	for results.Next() {
		err := results.Scan(&account.ID, &account.AccountName, &account.AccountNumber, &account.Balance)
		if err != nil {
			return model.Account{}
		}
	}
	return account

}
