package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var (
	connStr string = "user=limage dbname=thunder-prod sslmode=disable"
	db      *sql.DB
	err     error
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if os.Getenv("DATABASE_URL") != "" {
		connStr = os.Getenv("DATABASE_URL")
	}
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer db.Close()
	fmt.Println("kuru-")
	createAnnotationTable()
	createTransAnnotationTable()
	createTagTable()

	if os.Getenv("DATABASE_URL") == "" {
		copyAnnotations()
	}
	updateTransAnnotationsDB()

	router := httprouter.New()
	router.GET("/", Logging(Index, "index"))
	router.GET("/annotations", CommonHeaders(AnnotationIndex, "annotation-index"))
	router.GET("/annotations/:annotationId", IDShouldBeInt(AnnotationShow, "annotation-show", []string{"annotationId"}))

	router.GET("/select/annotations/:year/:month/:day/:hour/:min", IDShouldBeInt(SelectAnnotationIndex, "annotation-select-index", []string{"year", "month", "day", "hour", "min"}))
	router.GET("/select/area/annotations/:areaId/:year/:month/:day", IDShouldBeInt(SelectByAreaAnnotationIndex, "annotation-select-featured-index", []string{"areaId", "year", "month", "day"}))
	router.GET("/select/area/annotations/:areaId/:year/:month/:day/:featured", IDShouldBeInt(SelectByAreaAnnotationIndex, "annotation-select-featured-index", []string{"areaId", "year", "month", "day"}))

	router.GET("/trans/annotations", CommonHeaders(TransAnnotationIndex, "trans-annotation-index"))
	router.GET("/trans/annotations/:annotationId", IDShouldBeInt(TransAnnotationShow, "trans-annotation-show", []string{"annotationId"}))
	router.PUT("/trans/annotations/:annotationId", IDShouldBeInt(TransAnnotationUpdate, "trans-annotation-update", []string{"annotationId"}))

	router.GET("/tags", CommonHeaders(TagIndex, "tag-index"))
	router.POST("/tags", CommonHeaders(TagCreate, "tag-create"))
	router.DELETE("/tags/:tagId", IDShouldBeInt(TagDelete, "tag-delete", []string{"tagId"}))

	router.GET("/areas", CommonHeaders(AreaIndex, "area-index"))

	router.POST("/trans/annotations/:annotationId/niceNum/increment", IDShouldBeInt(IncrementNiceNum, "increment-niceNum", []string{"annotationId"}))
	router.POST("/trans/annotations/:annotationId/niceNum/decrement", IDShouldBeInt(DecrementNiceNum, "decrement-niceNum", []string{"annotationId"}))
	router.POST("/trans/annotations/:annotationId/updateFeatured/:featured", IDShouldBeInt(UpdateFeatured, "update-featured", []string{"annotationId"}))

	router.GET("/ranking", CommonHeaders(Ranking, "ranking"))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
