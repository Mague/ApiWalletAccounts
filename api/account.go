package api

import (
	"encoding/json"
	"fmt"
	"github.com/Mague/ApiWalletAccounts/models"
	"net/http"
)

func AllAccounts(res http.ResponseWriter, req *http.Request) {
	var account models.Account

	account.Id = "1"
	account.UserName = "Maguerencor"
	account.Email = "enmanueldavidmolina@gmail.com"
	account.Password = "enmanuel"
	account.WebSite = "i-say.com"

	jsonResult, err := json.Marshal(account)
	if err != nil {
		fmt.Fprint(res, "Error al generar el json")
		return
	}
	fmt.Println("All")
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonResult)
}
