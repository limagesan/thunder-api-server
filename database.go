package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func testTransaction() {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// テーブル作成
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "BOOKS" ("ID" INTEGER PRIMARY KEY, "TITLE" VARCHAR(255))`,
	)
	if err != nil {
		panic(err)
	}
}

func insertData() {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	// データの挿入
	res, err := db.Exec(
		`INSERT INTO BOOKS (ID, TITLE) VALUES (?, ?)`,
		123,
		"title",
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

func getLists() {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	// 複数レコード取得
	rows, err := db.Query(
		`SELECT * FROM BOOKS`,
	)
	if err != nil {
		panic(err)
	}

	// 処理が終わったらカーソルを閉じる
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string

		// カーソルから値を取得
		if err := rows.Scan(&id, &title); err != nil {
			log.Fatal("rows.Scan()", err)
			return
		}

		fmt.Printf("id: %d, title: %s\n", id, title)
	}
}

func getOne(id int) {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	// 1件取得
	row := db.QueryRow(
		`SELECT * FROM BOOKS WHERE ID=?`,
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

func updateData(id int) {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	// 更新
	res, err := db.Exec(
		`UPDATE BOOKS SET TITLE=? WHERE ID=?`,
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

func deleteData(id int) {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	// 削除
	res, err := db.Exec(
		`DELETE FROM BOOKS WHERE ID=?`,
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
