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

	// scraping()
	createTable()
	// insertTestData()
	annotations := getLists()
	RepoUpdateAnnotations(annotations)

	router := httprouter.New()
	router.GET("/", Logging(Index, "index"))
	router.GET("/annotations", CommonHeaders(AnnotationIndex, "annotation-index"))
	router.GET("/annotations/:annotationId", IDShouldBeInt(AnnotationShow, "annotation-show"))
	router.POST("/annotations", CommonHeaders(AnnotationCreate, "annotation-create"))
	router.DELETE("/annotations/:annotationId", IDShouldBeInt(AnnotationDelete, "annotation-delete"))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
