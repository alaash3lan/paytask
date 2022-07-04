package account

import (
	"encoding/json"
	"github.com/alaash3lan/paytask/tools"
	"io/ioutil"
	"log"
	"net/http"
)

// account data
type Account struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance,string"`
}

func GetAccounts(url string) map[string]*Account {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data []*Account
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	var accounts = make(map[string]*Account, 0)
	for _, value := range data {
		accounts[value.Id] = value
	}
	log.Print("finished ingesting accounts json data")
	return accounts
}

// transaction handler
func TransferHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transaction Transaction
	json.NewDecoder(r.Body).Decode(&transaction)
	err := transaction.Transfer(r.Context())
	if err != nil {
		json.NewEncoder(w).Encode(tools.Response{
			Status:  400,
			Success: false,
			Error:   err.Error(),
			Message: nil,
		})
		return
	}
	json.NewEncoder(w).Encode(tools.Response{
		Status:  200,
		Success: true,
		Error:   "",
		Message: "successful transaction",
	})
}
