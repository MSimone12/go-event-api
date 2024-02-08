package router

import (
	"go-event-api/src/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func JsonMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf8")
			next.ServeHTTP(w, req)
		})
	}
}

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(JsonMiddleware())
	router.HandleFunc("/reset", controller.HandleReset).Methods(http.MethodPost)
	router.HandleFunc("/balance", controller.HandleBalance).Methods(http.MethodGet)
	router.HandleFunc("/event", controller.HandleEvent).Methods(http.MethodPost)
	return router
}
