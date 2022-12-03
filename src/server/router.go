package server

import (
	"net/http"

	mux "github.com/gorilla/mux"
	"github.com/kalpit-sharma-dev/url-shortner/src/handler"
	"github.com/kalpit-sharma-dev/url-shortner/src/repository/db"
	"github.com/kalpit-sharma-dev/url-shortner/src/service"
)

func LoadRoute() {

	router := mux.NewRouter().PathPrefix("/url-shortner-service/api").Subrouter()
	router.Use(headerMiddleware)
	registerAppRoutes(router)
	http.ListenAndServe(":8080", router)
}

func registerAppRoutes(r *mux.Router) {
	urlRepo := db.NewFileRepository(db.GetFileProvider())
	urlService := service.NewUrlService(urlRepo)
	urlHandlers := handler.NewHandler(urlService)

	r.HandleFunc("/urls", urlHandlers.HandleGetUrl).Methods(http.MethodGet)
	r.HandleFunc("/urls", urlHandlers.HandleCreateUrl).Methods(http.MethodPost)
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
