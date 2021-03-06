package controllers

import (
	"github.com/mitchellh/hashstructure"
	"strconv"
)

type Event struct {
	Name string
	Date string
	Key string
}

func HashValue(c Event) uint64 {
	hash, err := hashstructure.Hash(c, nil)
	if err != nil {
		panic(err)
	}
	return hash
}

func AddEvent(name string, date string) Event {
	newEvent := Event{name, date, ""}
	newEvent.Key = strconv.FormatUint(HashValue(newEvent), 20)
	return newEvent
}