package dac

import "model"

type HistoryDac struct {
	ConnString string
}

func NewHistoryDac(conn string) *HistoryDac {
	return &HistoryDac{ConnString: conn}
}

func (historyDac *HistoryDac) ReadByAccountNumber(accountNumber string) []model.History {
	db := openDB(historyDac.ConnString)
	defer db.Close()
	results, _ := db.Query("SELECT * FROM history WHERE account_number=?", accountNumber)
	var histories []model.History
	for results.Next() {
		var history model.History
		err := results.Scan(&history.ID, &history.AccountNumber, &history.Description, &history.Success, &history.Date)
		if err != nil {
			return []model.History{}
		}
		histories = append(histories, history)
	}
	return histories

}

func (historyDac *HistoryDac) AddHistory(history model.History) bool {
	db := openDB(historyDac.ConnString)
	defer db.Close()
	results, _ := db.Prepare(`Insert INTO history(account_number,description,success,date) VALUES(?,?,?,?)`)

	_, err := results.Exec(history.AccountNumber, history.Description, history.Success, history.Date)
	if err != nil {
		return false
	}
	return true
}
