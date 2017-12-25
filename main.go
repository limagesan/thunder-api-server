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
	createAnnotationTable()
	createTagTable()

	// insertTestData()
	annotations := getAnnotations()
	RepoUpdateAnnotations(annotations)

	router := httprouter.New()
	router.GET("/", Logging(Index, "index"))
	router.GET("/annotations", CommonHeaders(AnnotationIndex, "annotation-index"))
	router.GET("/annotations/:annotationId", IDShouldBeInt(AnnotationShow, "annotation-show", "annotationId"))
	router.POST("/annotations", CommonHeaders(AnnotationCreate, "annotation-create"))
	router.DELETE("/annotations/:annotationId", IDShouldBeInt(AnnotationDelete, "annotation-delete", "annotationId"))

	router.GET("/tags", CommonHeaders(TagIndex, "tag-index"))
	router.POST("/tags", CommonHeaders(TagCreate, "tag-create"))
	router.DELETE("/tags/:tagId", IDShouldBeInt(TagDelete, "tag-delete", "tagId"))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
