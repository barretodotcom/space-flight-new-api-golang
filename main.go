package main

import (
	"log"
	"net/http"

	"articles/artc_http_func"

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
	r.HandleFunc("/articles", artc_http_func.GetArticles)
	http.ListenAndServe(":3000", r)

}
