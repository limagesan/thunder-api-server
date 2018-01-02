package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Annotationsに存在しているデータのTransAnnotationデータがなかったら追加する
func updateTransAnnotationsDB() {
	annotations := getAllAnnotations()

	var willDeleteIds []int
	var willAddIds []int
	transAnnotations := getTransAnnotations()

	for i := 0; i < len(annotations); i++ {
		var isFound = false
		for j := 0; j < len(transAnnotations); j++ {
			if annotations[i].ID == transAnnotations[j].ID {
				isFound = true
			}
		}
		if isFound == false {
			willAddIds = append(willAddIds, annotations[i].ID)
		}
	}

	for i := 0; i < len(transAnnotations); i++ {
		var isFound = false
		for j := 0; j < len(annotations); j++ {
			if transAnnotations[i].ID == annotations[j].ID {
				isFound = true
			}
		}
		if isFound == false {
			willDeleteIds = append(willDeleteIds, transAnnotations[i].ID)
		}
	}
	fmt.Println("willAddIds", willAddIds, "willDeleteIds", willDeleteIds)
	for i := 0; i < len(willAddIds); i++ {
		var tagIds []int

		annotation := NewTransAnnotation(willAddIds[i], tagIds, 0)
		insertTransAnnotation(*annotation)
	}

	for i := 0; i < len(willDeleteIds); i++ {
		deleteTransAnnotation(willDeleteIds[i])
	}
}

func createTransAnnotationTable() {
	_, err = db.Exec(
		`create table if not exists "transannotations" ("id" integer primary key unique, "tagids" text, "nicenum" integer)`,
	)
	if err != nil {
		panic(err)
	}
}

func insertTransAnnotation(transAnnotation TransAnnotation) {
	var id int
	err := db.QueryRow(
		`insert into transannotations (id, tagids, nicenum) values ($1,$2,$3) returning id`,
		transAnnotation.ID, intSliceToString(transAnnotation.TagIds), transAnnotation.NiceNum).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("lastInsertId", id)
}

func getTransAnnotations() TransAnnotations {
	rows, err := db.Query(`select * from transannotations`)
	if err != nil {
		panic(err)
	}
	var transAnnotations TransAnnotations
	// 処理が終わったらカーソルを閉じる
	defer rows.Close()
	for rows.Next() {

		var ID int
		var TagIds string
		var NiceNum int

		// カーソルから値を取得
		if err := rows.Scan(&ID, &TagIds, &NiceNum); err != nil {
			log.Fatal("rows.Scan()", err)
			return transAnnotations
		}
		transAnnotation := NewTransAnnotation(ID, stringToIntSlice(TagIds), NiceNum)
		transAnnotations = append(transAnnotations, *transAnnotation)

		fmt.Printf("ID: %d, TagIds: %s, NiceNum: %d", ID, TagIds, NiceNum)
	}
	return transAnnotations
}

func getTransAnnotation(id int) TransAnnotation {
	row := db.QueryRow(
		`select * from transannotations where id=$1`,
		id,
	)

	var id2 int
	var tagIds string
	var niceNum int
	err2 := row.Scan(&id2, &tagIds, &niceNum)

	// エラー内容による分岐
	switch {
	case err2 == sql.ErrNoRows:
		fmt.Printf("Not found")
	case err2 != nil:
		panic(err2)
	default:
		fmt.Printf("id: %d, tagIds: %s, niceNum: %d\n", id2, tagIds, niceNum)
	}

	transAnnotation := NewTransAnnotation(id2, stringToIntSlice(tagIds), niceNum)
	return *transAnnotation
}

func updateTransAnnotation(id int, transAnnotation TransAnnotation) []int {

	// 同じタグIDが複数存在したら一つにする
	var tempArr []int
	for i := 0; i < len(transAnnotation.TagIds); i++ {
		found := false
		for k := 0; k < len(tempArr); k++ {
			if transAnnotation.TagIds[i] == tempArr[k] {
				found = true
			}
		}
		if !found {
			tempArr = append(tempArr, transAnnotation.TagIds[i])
		}
	}
	transAnnotation.TagIds = tempArr

	// 更新
	res, err := db.Exec(
		`update transannotations set tagids=$1 where id=$2`,
		intSliceToString(transAnnotation.TagIds),
		id)
	if err != nil {
		panic(err)
	}

	// 更新されたレコード数
	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("affected by update: %d\n", affect)
	return transAnnotation.TagIds
}

func deleteTransAnnotation(id int) {
	res, err := db.Exec(
		`delete from transannotations where id=$1`,
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

func incrementNiceNum(id int) TransAnnotation {
	// 更新
	currentAnnotation := getTransAnnotation(id)
	currentAnnotation.NiceNum = currentAnnotation.NiceNum + 1

	res, err := db.Exec(
		`update transannotations set nicenum=$1 where id=$2`,
		currentAnnotation.NiceNum,
		id)
	if err != nil {
		panic(err)
	}

	// 更新されたレコード数
	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("affected by update: %d\n", affect)
	return currentAnnotation
}

func decrementNiceNum(id int) TransAnnotation {
	// 更新
	currentAnnotation := getTransAnnotation(id)
	if currentAnnotation.NiceNum > 0 {
		currentAnnotation.NiceNum = currentAnnotation.NiceNum - 1
	}

	res, err := db.Exec(
		`update transannotations set nicenum=$1 where id=$2`,
		currentAnnotation.NiceNum,
		id)
	if err != nil {
		panic(err)
	}

	// 更新されたレコード数
	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("affected by update: %d\n", affect)
	return currentAnnotation
}
