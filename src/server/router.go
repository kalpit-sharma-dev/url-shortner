package server

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

func LoadRoute() {

	router := mux.NewRouter().PathPrefix("/url-shortner-service/api").Subrouter()
	router.Use(headerMiddleware)
	registerAppRoutes(router)
	http.ListenAndServe(":8080", router)
}

func registerAppRoutes(r *mux.Router) {

}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
