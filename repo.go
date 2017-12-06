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
	RepoCreateEvent(Event{Title: "中村パーキングYOYOライブ", LocationName: "北浦和Kyara"})
	RepoCreateEvent(Event{Title: "Freebeeがやっちゃるけえ", LocationName: "北浦和エアーズ"})
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
