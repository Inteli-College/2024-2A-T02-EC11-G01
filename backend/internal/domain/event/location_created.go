package event

import "time"

type LocationCreated struct {
	Name    string
	Payload interface{}
}

func NewLocationCreated() *LocationCreated {
	return &LocationCreated{
		Name: "LocationCreated",
	}
}

func (e *LocationCreated) GetName() string {
	return e.Name
}

func (e *LocationCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *LocationCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *LocationCreated) GetDateTime() time.Time {
	return time.Now()
}
