package main

import (
	"github.com/Mague/ApiWalletAccounts/api"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/accounts", api.AllAccounts)
	http.ListenAndServe(":3000", mux)
}
