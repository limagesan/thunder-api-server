package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
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
		if !isFound {
			willDeleteIds = append(willDeleteIds, transAnnotations[i].ID)
		}
	}

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
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}

	// テーブル作成
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "TRANSANNOTATIONS" ("ID" INTEGER PRIMARY KEY, "TAGIDS" INTEGER, "NICENUM" INTEGER)`,
	)
	if err != nil {
		panic(err)
	}
}

func insertTransAnnotation(transAnnotation TransAnnotation) {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// データの挿入
	res, err := db.Exec(
		`INSERT INTO TRANSANNOTATIONS (ID, TAGIDS, NICENUM) VALUES (?,?,?)`,
		transAnnotation.ID, intSliceToString(transAnnotation.TagIds), transAnnotation.NiceNum)
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

func getTransAnnotations() TransAnnotations {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}

	// 複数レコード取得
	rows, err := db.Query(`SELECT * FROM TRANSANNOTATIONS`)
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
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// 1件取得
	row := db.QueryRow(
		`SELECT * FROM TRANSANNOTATIONS WHERE ID=?`,
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
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// 更新

	res, err := db.Exec(
		`UPDATE TRANSANNOTATIONS SET TAGIDS=? WHERE ID=?`,
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
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// 削除
	res, err := db.Exec(
		`DELETE FROM TRANSANNOTATIONS WHERE ID=?`,
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
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// 更新
	currentAnnotation := getTransAnnotation(id)
	currentAnnotation.NiceNum = currentAnnotation.NiceNum + 1

	res, err := db.Exec(
		`UPDATE TRANSANNOTATIONS SET NICENUM=? WHERE ID=?`,
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
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// 更新
	currentAnnotation := getTransAnnotation(id)
	if currentAnnotation.NiceNum > 0 {
		currentAnnotation.NiceNum = currentAnnotation.NiceNum - 1
	}

	res, err := db.Exec(
		`UPDATE TRANSANNOTATIONS SET NICENUM=? WHERE ID=?`,
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
