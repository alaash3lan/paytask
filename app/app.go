package app

import (
	"context"
	"fmt"
	"github.com/alaash3lan/paytask/account"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var accounts map[string]*account.Account

//start the routes
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
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "db", accounts)))
	})
}

func handle(r *mux.Router) {
	//get all accounts
	r.HandleFunc("/accounts", account.All).Methods("GET")

	//get account by an id
	r.HandleFunc("/accounts/{id}", account.Get).Methods("GET")

	// transfer balance
	r.HandleFunc("/transaction", account.TransferHandler).Methods("POST")
}
