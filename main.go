package main

import (
	"log"
	"net/http"

	"articles/artc_http_func"
	"articles/custom_crons"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", artc_http_func.Default).Methods("GET")
	r.HandleFunc("/articles", artc_http_func.GetArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", artc_http_func.FindOne).Methods("GET")
	r.HandleFunc("/articles", artc_http_func.InsertArticle).Methods("POST")
	r.HandleFunc("/articles", artc_http_func.PutArticle).Methods("PUT")
	r.HandleFunc("/articles/{id}", artc_http_func.DeleteArticle).Methods("DELETE")
	custom_crons.StartCron()
	http.ListenAndServe(":3000", r)

}
