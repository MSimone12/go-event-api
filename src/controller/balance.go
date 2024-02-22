package controller

import (
	"fmt"
	"net/http"
	"strconv"
)

// HandleBalance handles the request for GET /balance
func HandleBalance(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("account_id")

	accountId, _ := strconv.Atoi(id)

	balance, err := controller.UseCase.GetBalance(uint(accountId))
	if err != nil {
		http.Error(w, "0", 404)
		return
	}
	stringBalance := fmt.Sprintf("%d", balance)
	w.Write([]byte(stringBalance))
}
