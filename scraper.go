package main

import "time"

func scraping() {
	annotation := NewAnnotation("サイガーガールアルバムリリースライブ", "https://image1.jp", "https://image2.jp", "https://image3.jp", time.Now(), time.Now(), 2000, "https://source.com", "ヒマラヤパーク", 13.24242, 134.3242)
	insertData(*annotation)
	getLists()
}
