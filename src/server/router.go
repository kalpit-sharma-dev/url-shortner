package server

import (
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
	"github.com/kalpit-sharma-dev/url-shortner/src/handler"
	"github.com/kalpit-sharma-dev/url-shortner/src/repository/db"
	"github.com/kalpit-sharma-dev/url-shortner/src/service"
)

func LoadRoute() {
	log.Println("INFO : Loading Router")
	router := mux.NewRouter().PathPrefix("/url-shortner-service/api").Subrouter()
	router.Use(headerMiddleware)
	registerAppRoutes(router)
	log.Println("INFO : Router Loaded Successfully")
	log.Println("INFO : Application is started Successfully")
	http.ListenAndServe(":8080", router)
}

func registerAppRoutes(r *mux.Router) {
	log.Println("INFO : Registering Router ")
	urlRepo := db.NewFileRepository(db.GetFileProvider())
	urlService := service.NewUrlService(urlRepo)
	urlHandlers := handler.NewHandler(urlService)

	//r.HandleFunc("/urls/{url}", urlHandlers.HandleGetUrl).Methods(http.MethodGet)
	r.HandleFunc("/urls", urlHandlers.HandleCreateUrl).Methods(http.MethodPost)
	log.Println("INFO : Router Registered Successfully")
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
