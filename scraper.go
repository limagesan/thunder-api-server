package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func scraping() {
	doc, err := goquery.NewDocument("http://mu-seum.co.jp/schedule.html")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	doc.Find("table.listCal").Each(func(_ int, s *goquery.Selection) {
		var title string
		title = s.Find("p.live_title").Text()
		var artists string
		artists = s.Find("strong").Text()
		fmt.Println(title, artists)
	})

	// doc.Find("p.live_title").Each(func(_ int, s *goquery.Selection) {
	// 	// url, _ := s.Attr("href")
	// 	fmt.Println(s.Text())
	// })
}

func insertTestData() {
	// annotation := NewAnnotation(0, "サイガーガールアルバムリリースライブ", "https://lastfm-img2.akamaized.net/i/u/300x300/3aba9f2a172ea19d48323d4f8c600638.png", "", "", time.Now().String(), time.Now().String(), 2000, "http://wallwall.tokyo/schedule/rau-def-%E3%80%8Cunisex%E3%80%8Drelease-party-day-time-ver/", "ヒマラヤパーク", 35.872296, 139.646788)
	// insertData(*annotation)
}
