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
	router.GET("/todos", CommonHeaders(TodoIndex, "todo-index"))
	router.GET("/todos/:todoId", IDShouldBeInt(TodoShow, "todo-show"))
	router.POST("/todos", CommonHeaders(TodoCreate, "todo-create"))
	router.DELETE("/todos/:todoId", IDShouldBeInt(TodoDelete, "todo-delete"))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
