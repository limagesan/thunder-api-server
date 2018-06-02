package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func createTagTable() {
	// テーブル作成
	_, err = db.Exec(
		`create table if not exists "tags" ("id" serial primary key unique, "name" varchar(255), "color" varchar(255))`,
	)
	if err != nil {
		panic(err)
	}
}

func insertTag(tag Tag) int {
	// データの挿入
	var id int
	err := db.QueryRow(`insert into tags (name, color) values ($1,$2) returning id`, tag.Name, tag.Color).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("lastInsertId", id)
	return int(id)
}

func getAreas() Areas {
	// 複数レコード取得
	rows, err := db.Query(`select * from areas`)
	if err != nil {
		panic(err)
	}
	var areas Areas
	// 処理が終わったらカーソルを閉じる
	defer rows.Close()
	for rows.Next() {

		var ID int
		var Name string
		var NameJp string

		// カーソルから値を取得
		if err := rows.Scan(&ID, &Name, &NameJp); err != nil {
			log.Fatal("rows.Scan()", err)
			return areas
		}
		area := NewArea(ID, Name, NameJp)
		areas = append(areas, *area)

		fmt.Printf("ID: %d, Name: %s, NameJp: %s", ID, Name, NameJp)
	}
	return areas
}

func getTags() Tags {
	// 複数レコード取得
	rows, err := db.Query(`select * from tags`)
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
	// 削除
	res, err := db.Exec(
		`delete from tags where id=$1`,
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
