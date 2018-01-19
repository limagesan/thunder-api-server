package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const CACHEDATE = 5

func removeAllAnnotations() {
	_, err = db.Exec(`delete from annotations`)
	if err != nil {
		panic(err)
	}
}

func copyAnnotations() {
	_db, err := sql.Open("postgres", "user=hiroki dbname=thunder-scrape sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer _db.Close()
	rows, err := _db.Query(`select * from annotations`)
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
		addNewAnnotation(_annotations[i])
	}
}

func addNewAnnotation(annotation Annotation) {
	var id int
	rows, err := db.Query(
		`insert into "annotations" (title, artists, description, artistimageurls, locationimageurls, videoids, starttime, endtime, timetext, pricetext, sourceurls, locationname, latitude, longitude) select $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14 where not exists(select * from annotations where title = $15) returning id`,
		annotation.Title, sliceToString(annotation.Artists), annotation.Description, sliceToString(annotation.ArtistImageURLs), sliceToString(annotation.LocationImageURLs), sliceToString(annotation.VideoIds), annotation.StartTime, annotation.EndTime, annotation.TimeText, annotation.PriceText, sliceToString(annotation.SourceURLs), annotation.LocationName, annotation.Coordinate.Latitude, annotation.Coordinate.Longitude, annotation.Title)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println("lastInsertId", id)

	}
}

func createAnnotationTable() {
	_, err = db.Exec(
		`create table if not exists "annotations" ("id" serial primary key unique, "title" text, "artists" text, "description" text,"artistimageurls" text, "locationimageurls" text ,"videoids" text,"starttime" text, "endtime" text, "timetext" text, "pricetext" text, "sourceurls" text, "locationname" text, "latitude" float8, "longitude" float8)`,
	)
	if err != nil {
		panic(err)
	}
}

func insertAnnotation(annotation Annotation) {
	var id int
	err = db.QueryRow(
		`insert into "annotations" (title, artists, description, artistimageurls, locationimageurls, videoids, starttime, endtime, timetext, pricetext, sourceurls, locationname, latitude, longitude) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) returning id`,
		annotation.Title, sliceToString(annotation.Artists), annotation.Description, sliceToString(annotation.ArtistImageURLs), sliceToString(annotation.LocationImageURLs), sliceToString(annotation.VideoIds), annotation.StartTime, annotation.EndTime, annotation.TimeText, annotation.PriceText, sliceToString(annotation.SourceURLs), annotation.LocationName, annotation.Coordinate.Latitude, annotation.Coordinate.Longitude,
	).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("lastInsertId", id)
}

func getAllAnnotations() Annotations {
	// 複数レコード取得
	rows, err := db.Query(`select * from annotations order by	starttime`)
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

func getAnnotations(headTime time.Time, tailTime time.Time) FullAnnotations {
	rows, err := db.Query(
		`select annotations.id, title, artists, description, artistimageurls, locationimageurls, videoids, starttime, endtime, timetext, pricetext, sourceurls, locationname, latitude, longitude, tagids, nicenum from annotations inner join transannotations on annotations.id = transannotations.id where $1 < endtime and starttime < $2 order by starttime`,
		headTime.String(),
		tailTime.String(),
	)
	if err != nil {
		panic(err)
	}
	var annotations FullAnnotations
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
		var tagIds string
		var niceNum int

		// カーソルから値を取得
		if err := rows.Scan(&ID, &Title, &Artists, &Description, &ArtistImageURLs, &LocationImageURLs, &VideoIds, &StartTime, &EndTime, &TimeText, &PriceText, &SourceURLs, &LocationName, &Coordinate.Latitude, &Coordinate.Longitude, &tagIds, &niceNum); err != nil {
			log.Fatal("rows.Scan()", err)
			return annotations
		}
		annotation := NewFullAnnotation(ID, Title, stringToSlice(Artists), Description, stringToSlice(ArtistImageURLs), stringToSlice(LocationImageURLs), stringToSlice(VideoIds), StartTime, EndTime, TimeText, PriceText, stringToSlice(SourceURLs), LocationName, Coordinate.Latitude, Coordinate.Longitude, stringToIntSlice(tagIds), niceNum)

		annotations = append(annotations, *annotation)

		fmt.Printf("ID: %d, Title: %s, Artists: %s, Description: %s, ArtistImageURLs: %s, LocationImageURLs: %s, VideoIds: %s, StartTime: %s, EndTime: %s, TimeText: %s, PriceText: %s, SourceURLs: %s, LocationName: %s, Latitude: %f, Longitude: %f, tagIds: %s, niceMun: %d\n",
			ID, Title, Artists, Description, ArtistImageURLs, LocationImageURLs, VideoIds, StartTime, EndTime, TimeText, PriceText, SourceURLs, LocationName, Coordinate.Latitude, Coordinate.Longitude, tagIds, niceNum)
	}
	return annotations
}

func getAnnotation(id int) Annotation {
	var annotation *Annotation

	row := db.QueryRow(
		`select * from annotations where id=$1`,
		id,
	)

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
	if err := row.Scan(&ID, &Title, &Artists, &Description, &ArtistImageURLs, &LocationImageURLs, &VideoIds, &StartTime, &EndTime, &TimeText, &PriceText, &SourceURLs, &LocationName, &Coordinate.Latitude, &Coordinate.Longitude); err != nil {
		log.Fatal("rows.Scan()", err)
		return *annotation
	}

	// エラー内容による分岐
	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("Not found")
	case err != nil:
		panic(err)
	default:
	}

	annotation = NewAnnotation(ID, Title, stringToSlice(Artists), Description, stringToSlice(ArtistImageURLs), stringToSlice(LocationImageURLs), stringToSlice(VideoIds), StartTime, EndTime, TimeText, PriceText, stringToSlice(SourceURLs), LocationName, Coordinate.Latitude, Coordinate.Longitude)

	return *annotation
}

func updateAnnotation(id int) {
	res, err := db.Exec(
		`update annotations set title=$1 where id=$2`,
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
	res, err := db.Exec(
		`delete from annotations where id=$1`,
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
	now := time.Now()
	tailTime := now.Add(30 * 24 * time.Hour)
	sampleLimit := 30
	// 複数レコード取得
	rows, err := db.Query(
		`select annotations.id, title, artists, description, artistimageurls, locationimageurls, videoids, starttime, endtime, timetext, pricetext, sourceurls, locationname, latitude, longitude, tagids, nicenum from annotations inner join transannotations on annotations.id = transannotations.id where $1 < endtime and endtime < $2  order by nicenum desc, starttime asc limit $3`,
		now.String(),
		tailTime.String(),
		sampleLimit)
	if err != nil {
		panic(err)
	}

	// WHERE $ < STARTTIME AND STARTTIME < $ ORDER BY	STARTTIME`,
	// 	headTime.String(),
	// 	tailTime.String(),
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
