package main

type TransAnnotation struct {
	ID      int      `json:"id"`
	TagIds  []string `json:"tagIds"`
	NiceNum int      `json:"niceNum"`
}

type TransAnnotations []TransAnnotation

func NewTransAnnotation(ID int, TagIds []string, NiceNum int) *TransAnnotation {
	p := new(TransAnnotation)
	p.ID = ID
	p.TagIds = TagIds
	p.NiceNum = NiceNum
	return p
}
