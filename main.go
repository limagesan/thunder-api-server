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
	router.GET("/events", CommonHeaders(EventIndex, "event-index"))
	router.GET("/events/:eventId", IDShouldBeInt(EventShow, "event-show"))
	router.POST("/events", CommonHeaders(EventCreate, "event-create"))
	router.DELETE("/events/:eventId", IDShouldBeInt(EventDelete, "event-delete"))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
