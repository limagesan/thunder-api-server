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

	removeAllAnnotations()
	copyAnnotations()
	// insertTestData()

	annotations := getAnnotations()
	updateTransAnnotationsDB(annotations)
	RepoUpdateAnnotations(annotations)

	router := httprouter.New()
	router.GET("/", Logging(Index, "index"))
	router.GET("/annotations", CommonHeaders(AnnotationIndex, "annotation-index"))
	router.GET("/annotations/:annotationId", IDShouldBeInt(AnnotationShow, "annotation-show", "annotationId"))
	router.POST("/annotations", CommonHeaders(AnnotationCreate, "annotation-create"))
	router.DELETE("/annotations/:annotationId", IDShouldBeInt(AnnotationDelete, "annotation-delete", "annotationId"))

	router.GET("/trans/annotations", CommonHeaders(TransAnnotationIndex, "trans-annotation-index"))
	router.GET("/trans/annotations/:annotationId", IDShouldBeInt(TransAnnotationShow, "trans-annotation-show", "annotationId"))
	router.PUT("/trans/annotations/:annotationId", IDShouldBeInt(TransAnnotationUpdate, "trans-annotation-update", "annotationId"))

	router.GET("/tags", CommonHeaders(TagIndex, "tag-index"))
	router.POST("/tags", CommonHeaders(TagCreate, "tag-create"))
	router.DELETE("/tags/:tagId", IDShouldBeInt(TagDelete, "tag-delete", "tagId"))

	router.POST("/trans/annotations/:annotationId/niceNum/increment", IDShouldBeInt(IncrementNiceNum, "increment-niceNum", "annotationId"))
	router.POST("/trans/annotations/:annotationId/niceNum/decrement", IDShouldBeInt(DecrementNiceNum, "decrement-niceNum", "annotationId"))

	router.GET("/ranking", CommonHeaders(Ranking, "ranking"))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
