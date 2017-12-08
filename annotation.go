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

func NewAnnotation(Title string, TopImageURL1 string, TopImageURL2 string, TopImageURL3 string, OpenTime time.Time, CloseTime time.Time, Price int, SourceURL string, LocationName string, Latitude float64, Longitude float64) *Annotation {
	p := new(Annotation)
	p.Title = Title
	p.TopImageURLs[0] = TopImageURL1
	p.TopImageURLs[1] = TopImageURL2
	p.TopImageURLs[2] = TopImageURL3
	p.OpenTime = OpenTime
	p.CloseTime = CloseTime
	p.Price = Price
	p.SourceURL = SourceURL
	p.LocationName = LocationName
	p.Coordinate.Latitude = Latitude
	p.Coordinate.Longitude = Longitude

	return p
}
