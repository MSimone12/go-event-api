package controller

import (
	"go-event-api/db"
	"net/http"
)

func HandleReset(w http.ResponseWriter, r *http.Request) {
	db.Reset()

	w.Write([]byte("OK"))
}
