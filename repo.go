package main

import (
	"fmt"
	"time"
)

var (
	annotations Annotations
	currentID   int
)

func init() {
	RepoCreateAnnotation(Annotation{Title: "中村パーキングYOYOライブ", LocationName: "北浦和Kyara"})
	RepoCreateAnnotation(Annotation{Title: "Freebeeがやっちゃるけえ", LocationName: "北浦和エアーズ"})
}

func RepoFindAnnotation(id int) Annotation {
	for _, t := range annotations {
		if t.ID == id {
			return t
		}
	}
	return Annotation{}
}

func RepoUpdateAnnotations(arrays Annotations) {
	annotations = arrays
}

func RepoCreateAnnotation(e Annotation) Annotation {
	currentID += 1
	e.ID = currentID
	_openTime := time.Now()
	_closeTime := time.Now()
	e.CloseTime = _openTime.String()
	e.OpenTime = _closeTime.String()
	annotations = append(annotations, e)
	return e
}

func RepoDestroyAnnotation(id int) error {
	for i, t := range annotations {
		if t.ID == id {
			annotations = append(annotations[:i], annotations[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Annotation with id of %d to delete", id)
}
