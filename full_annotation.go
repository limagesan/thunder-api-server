package main

type FullAnnotation struct {
	ID                int        `json:"id"`
	Title             string     `json:"title"`
	Artists           []string   `json:"artists"`
	Tags              []string   `json:"tags"`
	Description       string     `json:"description"`
	ArtistImageURLs   []string   `json:"artistImageURLs"`
	LocationImageURLs []string   `json:"locationImageURLs"`
	VideoIds          []string   `json:"videoIds"`
	StartTime         string     `json:"startTime"`
	EndTime           string     `json:"endTime"`
	Price             int        `json:"price"`
	PriceText         string     `json:"priceText"`
	SourceURL         string     `json:"sourceURL"`
	LocationName      string     `json:"locationName"`
	Coordinate        Coordinate `json:"coordinate"`
	TagIds            []int      `json:"tagIds"`
	NiceNum           int        `json:"niceNum"`
}

type FullAnnotations []FullAnnotation

func NewFullAnnotation(ID int, Title string, Artists []string, Tags []string, Description string, ArtistImageURLs []string, LocationImageURLs []string, VideoIds []string, StartTime string, EndTime string, Price int, PriceText string, SourceURL string, LocationName string, Latitude float64, Longitude float64, TagIds []int, NiceNum int) *FullAnnotation {
	p := new(FullAnnotation)
	p.ID = ID
	p.Title = Title
	p.Artists = Artists
	p.Tags = Tags
	p.Description = Description
	p.ArtistImageURLs = ArtistImageURLs
	p.LocationImageURLs = LocationImageURLs
	p.VideoIds = VideoIds
	p.StartTime = StartTime
	p.EndTime = EndTime
	p.Price = Price
	p.PriceText = PriceText
	p.SourceURL = SourceURL
	p.LocationName = LocationName
	p.Coordinate.Latitude = Latitude
	p.Coordinate.Longitude = Longitude
	p.TagIds = TagIds
	p.NiceNum = NiceNum

	return p
}
