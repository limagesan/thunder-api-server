package main

type Annotation struct {
	ID              int        `json:"id"`
	Title           string     `json:"title"`
	Artists         []string   `json:"artists"`
	Tags            []string   `json:"tags"`
	Description     string     `json:"description"`
	ArtistImageURLs []string   `json:"artistImageURLs"`
	PlaceImageURLs  []string   `json:"placeImageURLs"`
	VideoIds        []string   `json:"videoIds"`
	StartTime       string     `json:"startTime"`
	EndTime         string     `json:"endTime"`
	Price           int        `json:"price"`
	PriceText       string     `json:"priceText"`
	SourceURL       string     `json:"sourceURL"`
	LocationName    string     `json:"locationName"`
	Coordinate      Coordinate `json:"coordinate"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Annotations []Annotation

func NewAnnotation(ID int, Title string, Artists []string, Tags []string, Description string, ArtistImageURLs []string, PlaceImageURLs []string, VideoIds []string, StartTime string, EndTime string, Price int, PriceText string, SourceURL string, LocationName string, Latitude float64, Longitude float64) *Annotation {
	p := new(Annotation)
	p.ID = ID
	p.Title = Title
	p.Artists = Artists
	p.Tags = Tags
	p.Description = Description
	p.ArtistImageURLs = ArtistImageURLs
	p.PlaceImageURLs = PlaceImageURLs
	p.VideoIds = VideoIds
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
