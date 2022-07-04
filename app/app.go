package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alaash3lan/paytask/account"
	"github.com/alaash3lan/paytask/tools"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var accounts map[string]*account.Account

func Start() {
	accounts = account.GetAccounts(os.Getenv("ACCOUNTS_URL"))

	r := mux.NewRouter()
	r.Use(withDB)

	handle(r)

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// add accounts data to the contexts
func withDB(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "db", accounts)))
	})
}

func handle(r *mux.Router) {
	//get all accounts
	r.HandleFunc("/accounts", func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(accounts)
	}).Methods("GET")

	//get account by an id
	r.HandleFunc("/accounts/{id}", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		id := mux.Vars(request)["id"]
		if accounts[id] == nil {
			json.NewEncoder(writer).Encode(tools.Response{
				Status:  404,
				Success: false,
				Error:   "wrong account id",
				Message: nil,
			})
		}
		json.NewEncoder(writer).Encode(accounts[id])
	}).Methods("GET")

	//
	r.HandleFunc("/transaction", account.TransferHandler).Methods("POST")
}
