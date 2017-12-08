package main

import "time"

type Annotation struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	TopImageURLs [3]string  `json:"topImageURLs"`
	OpenTime     time.Time  `json:"openTime"`
	CloseTime    time.Time  `json:"closeTime"`
	Price        int        `json:"price"`
	SourceURL    string     `json:"sourceURL"`
	LocationName string     `json:"locationName"`
	Coordinate   Coordinate `json:"coordinate"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Annotations []Annotation
