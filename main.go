package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := httprouter.New()
	router.GET("/", Logging(Index, "index"))
	router.GET("/todos", Logging(TodoIndex, "todo-index"))
	router.GET("/todos/:todoId", Logging(TodoShow, "todo-show"))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
