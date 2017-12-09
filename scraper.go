package main

import "time"

func scraping() {

}

func insertTestData() {
	annotation := NewAnnotation(0, "サイガーガールアルバムリリースライブ", "https://lastfm-img2.akamaized.net/i/u/300x300/3aba9f2a172ea19d48323d4f8c600638.png", "", "", time.Now().String(), time.Now().String(), 2000, "http://wallwall.tokyo/schedule/rau-def-%E3%80%8Cunisex%E3%80%8Drelease-party-day-time-ver/", "ヒマラヤパーク", 35.872296, 139.646788)
	insertData(*annotation)
}
