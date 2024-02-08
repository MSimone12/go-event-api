package controller

import (
	"go-event-api/src/db"
	"net/http"
)

func HandleReset(w http.ResponseWriter, r *http.Request) {
	db.Reset()

	w.Write([]byte("OK"))
}
