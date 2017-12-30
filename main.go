package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// db, err := sql.Open("sqlite3", "./database/trans-thunder.db")

	// res, err = db.Exec(`SELECT ID DESCRIPTION WHERE LOCATIONNAME=?`, "Contact Tokyo")

	// db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	// if err != nil {
	// 	panic(err)
	// }
	// // 削除
	// _, err = db.Exec(
	// 	`DROP TABLE ANNOTATIONS`,
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// return

	// scraping()
	createAnnotationTable()
	createTransAnnotationTable()
	createTagTable()

	// removeAllAnnotations()
	// copyAnnotations()
	// insertTestData()
	now := time.Now()
	annotations := getAnnotations(now, now)
	updateTransAnnotationsDB(annotations)
	// RepoUpdateAnnotations(annotations)

	router := httprouter.New()
	router.GET("/", Logging(Index, "index"))
	router.GET("/annotations", CommonHeaders(AnnotationIndex, "annotation-index"))

	router.GET("/annotations/:annotationId", IDShouldBeInt(AnnotationShow, "annotation-show", []string{"annotationId"}))
	router.POST("/annotations", CommonHeaders(AnnotationCreate, "annotation-create"))
	router.DELETE("/annotations/:annotationId", IDShouldBeInt(AnnotationDelete, "annotation-delete", []string{"annotationId"}))
	router.GET("/select/annotations/:year/:month/:day/:hour/:min", IDShouldBeInt(AnnotationIndex2, "annotation-select-index", []string{"year", "month", "day", "hour", "min"}))

	router.GET("/trans/annotations", CommonHeaders(TransAnnotationIndex, "trans-annotation-index"))
	router.GET("/trans/annotations/:annotationId", IDShouldBeInt(TransAnnotationShow, "trans-annotation-show", []string{"annotationId"}))
	router.PUT("/trans/annotations/:annotationId", IDShouldBeInt(TransAnnotationUpdate, "trans-annotation-update", []string{"annotationId"}))

	router.GET("/tags", CommonHeaders(TagIndex, "tag-index"))
	router.POST("/tags", CommonHeaders(TagCreate, "tag-create"))
	router.DELETE("/tags/:tagId", IDShouldBeInt(TagDelete, "tag-delete", []string{"tagId"}))

	router.POST("/trans/annotations/:annotationId/niceNum/increment", IDShouldBeInt(IncrementNiceNum, "increment-niceNum", []string{"annotationId"}))
	router.POST("/trans/annotations/:annotationId/niceNum/decrement", IDShouldBeInt(DecrementNiceNum, "decrement-niceNum", []string{"annotationId"}))

	router.GET("/ranking", CommonHeaders(Ranking, "ranking"))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
