package controller

import (
	"encoding/json"
	"fmt"
	"go-event-api/src/domain/entity"
	"go-event-api/src/domain/usecase"
	"net/http"
)

// HandleEvent handles the request for POST /event
func HandleEvent(w http.ResponseWriter, r *http.Request) {
	var event entity.Event
	json.NewDecoder(r.Body).Decode(&event)
	fmt.Println(event)
	switch event.Type {
	case "deposit":
		deposit(w, event)

	case "withdraw":
		withdraw(w, event)

	case "transfer":
		transfer(w, event)

	default:
		http.Error(w, "Unknow event type", 404)

	}
}

func deposit(w http.ResponseWriter, event entity.Event) {
	deposit, err := usecase.Deposit(event.GetDestination(), event.Amount)

	if err != nil {
		http.Error(w, "0", 404)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(deposit)
}

func withdraw(w http.ResponseWriter, event entity.Event) {
	withdraw, err := usecase.Withdraw(event.GetOrigin(), event.Amount)

	if err != nil {
		http.Error(w, "0", 404)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(withdraw)
}

func transfer(w http.ResponseWriter, event entity.Event) {
	transfer, err := usecase.Transfer(event.GetOrigin(), event.GetDestination(), event.Amount)

	if err != nil {
		http.Error(w, "0", 404)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(transfer)
}
