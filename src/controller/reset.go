package controller

import (
	"go-event-api/src/db"
	"net/http"
)

// HandleReset is an IO operation which should delete and recreate db.sqlite
func HandleReset(w http.ResponseWriter, r *http.Request) {
	db.Reset(db.DefaultDatabaseName)

	w.Write([]byte("OK"))
}
