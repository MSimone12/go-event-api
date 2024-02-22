package controller

import (
	"encoding/json"
	"go-event-api/src/domain/entity"
	"net/http"
)

// HandleEvent handles the request for POST /event,
// redirecting to a specific function depending on the body.type parameter
func HandleEvent(w http.ResponseWriter, r *http.Request) {
	var event entity.Event
	json.NewDecoder(r.Body).Decode(&event)
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

// deposit will handle the deposit event type, calling the UseCase.Deposit function
func deposit(w http.ResponseWriter, event entity.Event) {
	deposit := controller.UseCase.Deposit(event.GetDestination(), event.Amount)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(deposit)
}

// withdraw will handle the withdraw event type, calling the UseCase.Withdraw function
func withdraw(w http.ResponseWriter, event entity.Event) {
	withdraw, err := controller.UseCase.Withdraw(event.GetOrigin(), event.Amount)

	if err != nil {
		http.Error(w, "0", 404)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(withdraw)
}

// transfer will handle the transfer event type, calling the UseCase.Transfer function
func transfer(w http.ResponseWriter, event entity.Event) {
	transfer, err := controller.UseCase.Transfer(event.GetOrigin(), event.GetDestination(), event.Amount)

	if err != nil {
		http.Error(w, "0", 404)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(transfer)
}
