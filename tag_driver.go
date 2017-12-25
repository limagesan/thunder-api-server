package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createTagTable() {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}

	// テーブル作成
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "TAGS" ("ID" INTEGER PRIMARY KEY AUTOINCREMENT, "NAME" VARCHAR(255), "COLOR" VARCHAR(255))`,
	)
	if err != nil {
		panic(err)
	}
}

func insertTag(tag Tag) int {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// データの挿入
	res, err := db.Exec(`INSERT INTO TAGS (NAME, COLOR) VALUES (?,?)`, tag.Name, tag.Color)
	if err != nil {
		panic(err)
	}

	// 挿入処理の結果からIDを取得
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("lastInsertId", id)
	return int(id)
}

func getTags() Tags {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}

	// 複数レコード取得
	rows, err := db.Query(`SELECT * FROM TAGS`)
	if err != nil {
		panic(err)
	}
	var tags Tags
	// 処理が終わったらカーソルを閉じる
	defer rows.Close()
	for rows.Next() {

		var ID int
		var Name string
		var Color string

		// カーソルから値を取得
		if err := rows.Scan(&ID, &Name, &Color); err != nil {
			log.Fatal("rows.Scan()", err)
			return tags
		}
		tag := NewTag(ID, Name, Color)
		tags = append(tags, *tag)

		fmt.Printf("ID: %d, Name: %s, Color: %s", ID, Name, Color)
	}
	return tags
}

func deleteTag(id int) {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./database/trans-thunder.db")
	if err != nil {
		panic(err)
	}
	// 削除
	res, err := db.Exec(
		`DELETE FROM TAGS WHERE ID=?`,
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
