package main

import "time"

type Event struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	OpenTime    time.Time  `json:"openTime"`
	CloseTime   time.Time  `json:"closeTime"`
	Description string     `json:"description"`
	ImageURL    string     `json:"imageUrl"`
	Coordinate  Coordinate `json:"coordinate"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Events []Event
