package main

type FullAnnotation struct {
	ID                int        `json:"id"`
	Title             string     `json:"title"`
	Artists           []string   `json:"artists"`
	Description       string     `json:"description"`
	ArtistImageURLs   []string   `json:"artistImageURLs"`
	LocationImageURLs []string   `json:"locationImageURLs"`
	VideoIds          []string   `json:"videoIds"`
	StartTime         string     `json:"startTime"`
	EndTime           string     `json:"endTime"`
	TimeText          string     `json:"timeText"`
	PriceText         string     `json:"priceText"`
	SourceURLs        []string   `json:"sourceURLs"`
	LocationName      string     `json:"locationName"`
	Coordinate        Coordinate `json:"coordinate"`
	AreaId            int        `json:"areaId"`
	TagIds            []int      `json:"tagIds"`
	NiceNum           int        `json:"niceNum"`
}

type FullAnnotations []FullAnnotation

func NewFullAnnotation(ID int, Title string, Artists []string, Description string, ArtistImageURLs []string, LocationImageURLs []string, VideoIds []string, StartTime string, EndTime string, TimeText string, PriceText string, SourceURLs []string, LocationName string, Latitude float64, Longitude float64, AreaId int, TagIds []int, NiceNum int) *FullAnnotation {
	p := new(FullAnnotation)
	p.ID = ID
	p.Title = Title
	p.Artists = Artists
	p.Description = Description
	p.ArtistImageURLs = ArtistImageURLs
	p.LocationImageURLs = LocationImageURLs
	p.VideoIds = VideoIds
	p.StartTime = StartTime
	p.EndTime = EndTime
	p.TimeText = TimeText
	p.PriceText = PriceText
	p.SourceURLs = SourceURLs
	p.LocationName = LocationName
	p.Coordinate.Latitude = Latitude
	p.Coordinate.Longitude = Longitude
	p.AreaId = AreaId
	p.TagIds = TagIds
	p.NiceNum = NiceNum

	return p
}
