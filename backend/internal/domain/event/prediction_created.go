package event

import "time"

type PredictionCreated struct {
	Name    string
	Payload interface{}
}

func NewPredictionCreated() *PredictionCreated {
	return &PredictionCreated{
		Name: "PredictionCreated",
	}
}

func (e *PredictionCreated) GetName() string {
	return e.Name
}

func (e *PredictionCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *PredictionCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *PredictionCreated) GetDateTime() time.Time {
	return time.Now()
}
