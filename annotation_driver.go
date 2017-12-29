package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const CACHEDATE = 5

func removeAllAnnotations() {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// データの挿入
	_, err = db.Exec(`DELETE FROM ANNOTATIONS`)
	if err != nil {
		panic(err)
	}
}

func copyAnnotations() {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/thunder.db")
	if err != nil {
		panic(err)
	}

	// 複数レコード取得
	rows, err := db.Query(`SELECT * FROM ANNOTATIONS`)
	if err != nil {
		panic(err)
	}
	var _annotations Annotations
	// 処理が終わったらカーソルを閉じる
	defer rows.Close()
	for rows.Next() {

		var ID int
		var Title string
		var Artists string
		var Description string
		var ArtistImageURLs string
		var LocationImageURLs string
		var VideoIds string
		var StartTime string // Time.timeだとScan時にエラーになる
		var EndTime string
		var TimeText string
		var PriceText string
		var SourceURLs string
		var LocationName string
		var Coordinate Coordinate

		// カーソルから値を取得
		if err := rows.Scan(&ID, &Title, &Artists, &Description, &ArtistImageURLs, &LocationImageURLs, &VideoIds, &StartTime, &EndTime, &TimeText, &PriceText, &SourceURLs, &LocationName, &Coordinate.Latitude, &Coordinate.Longitude); err != nil {
			log.Fatal("rows.Scan()", err)
		}
		annotation := NewAnnotation(ID, Title, stringToSlice(Artists), Description, stringToSlice(ArtistImageURLs), stringToSlice(LocationImageURLs), stringToSlice(VideoIds), StartTime, EndTime, TimeText, PriceText, stringToSlice(SourceURLs), LocationName, Coordinate.Latitude, Coordinate.Longitude)

		_annotations = append(_annotations, *annotation)

		fmt.Printf("ID: %d, Title: %s, Artists: %s, Description: %s, ArtistImageURLs: %s, LocationImageURLs: %s, VideoIds: %s, StartTime: %s, EndTime: %s, TimeText: %s, PriceText: %s, SourceURL: %s, LocationName: %s, Latitude: %f, Longitude: %f\n",
			ID, Title, Artists, Description, ArtistImageURLs, LocationImageURLs, VideoIds, StartTime, EndTime, TimeText, PriceText, SourceURLs, LocationName, Coordinate.Latitude, Coordinate.Longitude)

	}

	fmt.Println("checkAnnotations", _annotations)
	for i := 0; i < len(_annotations); i++ {
		insertAnnotation(_annotations[i])
	}
}

func createAnnotationTable() {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}

	// テーブル作成
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "ANNOTATIONS" ("ID" INTEGER PRIMARY KEY, "TITLE" VARCHAR(255), "ARTISTS" VARCHAR(255), "DESCRIPTION" VARCHAR(255),"ARTISTIMAGEURLS" TEXT, "LOCATIONIMAGEURLS" TEXT ,"VIDEOIDS" TEXT,"STARTTIME" VARCHAR(255), "ENDTIME" VARCHAR(255), "TIMETEXT" VARCHAR(255), "PRICETEXT" VARCHAR(255), "SOURCEURLS" VARCHAR(255), "LOCATIONNAME" VARCHAR(255), "LATITUDE" REAL, "LONGITUDE" REAL)`,
	)
	if err != nil {
		panic(err)
	}
}

func insertAnnotation(annotation Annotation) {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// データの挿入
	res, err := db.Exec(
		`INSERT INTO ANNOTATIONS (ID, TITLE, ARTISTS, DESCRIPTION, ARTISTIMAGEURLS, LOCATIONIMAGEURLS, VIDEOIDS, STARTTIME, ENDTIME, TIMETEXT, PRICETEXT, SOURCEURLS, LOCATIONNAME, LATITUDE, LONGITUDE) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		annotation.ID, annotation.Title, sliceToString(annotation.Artists), annotation.Description, sliceToString(annotation.ArtistImageURLs), sliceToString(annotation.LocationImageURLs), sliceToString(annotation.VideoIds), annotation.StartTime, annotation.EndTime, annotation.TimeText, annotation.PriceText, sliceToString(annotation.SourceURLs), annotation.LocationName, annotation.Coordinate.Latitude, annotation.Coordinate.Longitude,
	)
	if err != nil {
		panic(err)
	}

	// 挿入処理の結果からIDを取得
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("lastInsertId", id)
}

func getAnnotations() Annotations {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	now := time.Now()
	afterTwoDays := now.Add(CACHEDATE * 24 * time.Hour)

	// 複数レコード取得
	rows, err := db.Query(
		`SELECT * FROM ANNOTATIONS WHERE ? < STARTTIME AND STARTTIME < ? ORDER BY	STARTTIME`,
		now.String(),
		afterTwoDays.String(),
	)
	if err != nil {
		panic(err)
	}
	var annotations Annotations
	// 処理が終わったらカーソルを閉じる
	defer rows.Close()
	for rows.Next() {

		var ID int
		var Title string
		var Artists string
		var Description string
		var ArtistImageURLs string
		var LocationImageURLs string
		var VideoIds string
		var StartTime string // Time.timeだとScan時にエラーになる
		var EndTime string
		var TimeText string
		var PriceText string
		var SourceURLs string
		var LocationName string
		var Coordinate Coordinate

		// カーソルから値を取得
		if err := rows.Scan(&ID, &Title, &Artists, &Description, &ArtistImageURLs, &LocationImageURLs, &VideoIds, &StartTime, &EndTime, &TimeText, &PriceText, &SourceURLs, &LocationName, &Coordinate.Latitude, &Coordinate.Longitude); err != nil {
			log.Fatal("rows.Scan()", err)
			return annotations
		}
		annotation := NewAnnotation(ID, Title, stringToSlice(Artists), Description, stringToSlice(ArtistImageURLs), stringToSlice(LocationImageURLs), stringToSlice(VideoIds), StartTime, EndTime, TimeText, PriceText, stringToSlice(SourceURLs), LocationName, Coordinate.Latitude, Coordinate.Longitude)

		annotations = append(annotations, *annotation)

		fmt.Printf("ID: %d, Title: %s, Artists: %s, Description: %s, ArtistImageURLs: %s, LocationImageURLs: %s, VideoIds: %s, StartTime: %s, EndTime: %s, TimeText: %s, PriceText: %s, SourceURLs: %s, LocationName: %s, Latitude: %f, Longitude: %f\n",
			ID, Title, Artists, Description, ArtistImageURLs, LocationImageURLs, VideoIds, StartTime, EndTime, TimeText, PriceText, SourceURLs, LocationName, Coordinate.Latitude, Coordinate.Longitude)
	}
	return annotations
}

func getAnnotation(id int) {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// 1件取得
	row := db.QueryRow(
		`SELECT * FROM ANNOTATIONS WHERE ID=?`,
		id,
	)

	var id2 int
	var title string
	err2 := row.Scan(&id2, &title)

	// エラー内容による分岐
	switch {
	case err2 == sql.ErrNoRows:
		fmt.Printf("Not found")
	case err2 != nil:
		panic(err2)
	default:
		fmt.Printf("id: %d, title: %s\n", id2, title)
	}
}

func updateAnnotation(id int) {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// 更新
	res, err := db.Exec(
		`UPDATE ANNOTATIONS SET TITLE=? WHERE ID=?`,
		"update title",
		id,
	)
	if err != nil {
		panic(err)
	}

	// 更新されたレコード数
	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("affected by update: %d\n", affect)
}

func deleteAnnotation(id int) {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// 削除
	res, err := db.Exec(
		`DELETE FROM ANNOTATIONS WHERE ID=?`,
		id,
	)
	if err != nil {
		panic(err)
	}

	// 削除されたレコード数
	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("affected by delete: %d\n", affect)
}

func sliceToString(sli []string) string {
	var S string = ""
	for i := 0; i < len(sli); i++ {
		var comma string
		if i == len(sli)-1 {
			comma = ""
		} else {
			comma = ","
		}
		S = S + sli[i] + comma
	}
	return S
}

func stringToSlice(str string) []string {
	rep := regexp.MustCompile(`\s*,\s*`)
	sep := rep.Split(str, -1)
	return sep
}

func intSliceToString(sli []int) string {
	var S string = ""
	for i := 0; i < len(sli); i++ {
		var comma string
		if i == len(sli)-1 {
			comma = ""
		} else {
			comma = ","
		}
		el := strconv.Itoa(sli[i])
		S = S + el + comma
	}
	return S
}

func stringToIntSlice(str string) []int {
	rep := regexp.MustCompile(`\s*,\s*`)
	sep := rep.Split(str, -1)
	var res []int
	for i := 0; i < len(sep); i++ {
		el, err := strconv.Atoi(sep[i])
		if err == nil {
			res = append(res, el)
		}
	}
	return res
}

func getRanking() FullAnnotations {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}

	// 複数レコード取得
	rows, err := db.Query(
		`SELECT ANNOTATIONS.ID, TITLE, ARTISTS, DESCRIPTION, ARTISTIMAGEURLS, LOCATIONIMAGEURLS, VIDEOIDS, STARTTIME, ENDTIME, TIMETEXT, PRICETEXT, SOURCEURLS, LOCATIONNAME, LATITUDE, LONGITUDE, TAGIDS, NICENUM FROM ANNOTATIONS INNER JOIN TRANSANNOTATIONS ON ANNOTATIONS.ID = TRANSANNOTATIONS.ID ORDER BY NICENUM DESC`)
	if err != nil {
		panic(err)
	}
	var fullAnnotations FullAnnotations
	// 処理が終わったらカーソルを閉じる
	defer rows.Close()
	for rows.Next() {

		var ID int
		var Title string
		var Artists string
		var Description string
		var ArtistImageURLs string
		var LocationImageURLs string
		var VideoIds string
		var StartTime string // Time.timeだとScan時にエラーになる
		var EndTime string
		var TimeText string
		var PriceText string
		var SourceURLs string
		var LocationName string
		var Coordinate Coordinate
		var TagIds string
		var NiceNum int

		// カーソルから値を取得
		if err := rows.Scan(&ID, &Title, &Artists, &Description, &ArtistImageURLs, &LocationImageURLs, &VideoIds, &StartTime, &EndTime, &TimeText, &PriceText, &SourceURLs, &LocationName, &Coordinate.Latitude, &Coordinate.Longitude, &TagIds, &NiceNum); err != nil {
			log.Fatal("rows.Scan()", err)
			return fullAnnotations
		}
		fullAnnotation := NewFullAnnotation(ID, Title, stringToSlice(Artists), Description, stringToSlice(ArtistImageURLs), stringToSlice(LocationImageURLs), stringToSlice(VideoIds), StartTime, EndTime, TimeText, PriceText, stringToSlice(SourceURLs), LocationName, Coordinate.Latitude, Coordinate.Longitude, stringToIntSlice(TagIds), NiceNum)

		fullAnnotations = append(fullAnnotations, *fullAnnotation)

		fmt.Printf("ID: %d, Title: %s, Artists: %s, Description: %s, ArtistImageURLs: %s, LocationImageURLs: %s, VideoIds: %s, StartTime: %s, EndTime: %s, TimeText: %s, PriceText: %s, SourceURLs: %s, LocationName: %s, Latitude: %f, Longitude: %f, TagIds: %s, NiceNum: %d \n",
			ID, Title, Artists, Description, ArtistImageURLs, LocationImageURLs, VideoIds, StartTime, EndTime, TimeText, PriceText, SourceURLs, LocationName, Coordinate.Latitude, Coordinate.Longitude, TagIds, NiceNum)
	}
	return fullAnnotations
}
