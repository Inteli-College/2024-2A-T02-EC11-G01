package main

import "github.com/google/wire"

// TODO: Refactor all of this

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
