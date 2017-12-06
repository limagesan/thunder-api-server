package main

import (
	"fmt"
	"time"
)

var (
	events    Events
	currentID int
)

func init() {
	RepoCreateEvent(Event{Title: "Write presentation"})
	RepoCreateEvent(Event{Title: "Host meetup"})
}

func RepoFindEvent(id int) Event {
	for _, t := range events {
		if t.ID == id {
			return t
		}
	}
	return Event{}
}

func RepoCreateEvent(e Event) Event {
	currentID += 1
	e.ID = currentID
	e.CloseTime = time.Now()
	e.OpenTime = time.Now()
	e.Description = "結成1周年に作成したアルバムのリリース記念ライブです"
	e.Coordinate = Coordinate{Latitude: 35.871236, Longitude: 139.6427601}
	e.ImageURL = "https://i.ytimg.com/vi/mTpczxY5r8o/maxresdefault.jpg"
	events = append(events, e)
	return e
}

func RepoDestroyEvent(id int) error {
	for i, t := range events {
		if t.ID == id {
			events = append(events[:i], events[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Event with id of %d to delete", id)
}
