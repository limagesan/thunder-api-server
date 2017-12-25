package main

type TransAnnotation struct {
	ID      int   `json:"id"`
	TagIds  []int `json:"tagIds"`
	NiceNum int   `json:"niceNum"`
}

type TransAnnotations []TransAnnotation

func NewTransAnnotation(ID int, TagIds []int, NiceNum int) *TransAnnotation {
	p := new(TransAnnotation)
	p.ID = ID
	p.TagIds = TagIds
	p.NiceNum = NiceNum
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
