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

type App struct {
	Router *mux.Router
	DB     map[string]*account.Account
}

func (a *App) Initialize() {

	a.Router = mux.NewRouter()
	a.DB = account.GetAccounts(os.Getenv("ACCOUNTS_URL"))
	a.Router.Use(a.withDB)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	//get all accounts
	a.Router.HandleFunc("/accounts", account.All).Methods("GET")

	//get account by an id
	a.Router.HandleFunc("/accounts/{id}", account.Get).Methods("GET")

	// transfer balance
	a.Router.HandleFunc("/transaction", account.TransferHandler).Methods("POST")

}

//Add app  to the contexts
func (a *App) withDB(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "DB", a.DB)))
	})
}

//Run the server
func (a *App) Run(addr string) {
	fmt.Printf("Starting server at port %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
