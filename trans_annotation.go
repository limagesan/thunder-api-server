package main

type TransAnnotation struct {
	ID       int   `json:"id"`
	TagIds   []int `json:"tagIds"`
	NiceNum  int   `json:"niceNum"`
	Featured bool  `json:"featured"`
}

type TransAnnotations []TransAnnotation

func NewTransAnnotation(ID int, TagIds []int, NiceNum int, Featured bool) *TransAnnotation {
	p := new(TransAnnotation)
	p.ID = ID
	p.TagIds = TagIds
	p.NiceNum = NiceNum
	p.Featured = Featured
	return p
}

type Tag struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Tags []Tag

func NewTag(ID int, Name string, Color string) *Tag {
	p := new(Tag)
	p.ID = ID
	p.Name = Name
	p.Color = Color
	return p
}

type Area struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	NameJp string `json:"name_jp"`
}

type Areas []Area

func NewArea(ID int, Name string, NameJp string) *Area {
	p := new(Area)
	p.ID = ID
	p.Name = Name
	p.NameJp = NameJp
	return p
}
