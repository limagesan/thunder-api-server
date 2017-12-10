package main

type Annotation struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	Artists      string     `json:"artists"`
	Description  string     `json:"description"`
	TopImageURLs [3]string  `json:"topImageURLs"`
	StartTime    string     `json:"startTime"`
	EndTime      string     `json:"endTime"`
	Price        int        `json:"price"`
	PriceText    string     `json:"priceText"`
	SourceURL    string     `json:"sourceURL"`
	LocationName string     `json:"locationName"`
	Coordinate   Coordinate `json:"coordinate"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Annotations []Annotation

func NewAnnotation(ID int, Title string, Artists string, Description string, TopImageURL1 string, TopImageURL2 string, TopImageURL3 string, StartTime string, EndTime string, Price int, PriceText string, SourceURL string, LocationName string, Latitude float64, Longitude float64) *Annotation {
	p := new(Annotation)
	p.ID = ID
	p.Title = Title
	p.Artists = Artists
	p.Description = Description
	p.TopImageURLs[0] = TopImageURL1
	p.TopImageURLs[1] = TopImageURL2
	p.TopImageURLs[2] = TopImageURL3
	p.StartTime = StartTime
	p.EndTime = EndTime
	p.Price = Price
	p.PriceText = PriceText
	p.SourceURL = SourceURL
	p.LocationName = LocationName
	p.Coordinate.Latitude = Latitude
	p.Coordinate.Longitude = Longitude

	return p
}
