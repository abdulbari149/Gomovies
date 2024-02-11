package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abdulbari149/gomovies/middlewares"
	"github.com/abdulbari149/gomovies/movies"
	"github.com/gorilla/mux"
)

func main() {

	movies.Init()

	r := mux.NewRouter()

	r.Use(middlewares.LoggingMiddleware)

	movies.InitHandlers(r)
	fmt.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
}
