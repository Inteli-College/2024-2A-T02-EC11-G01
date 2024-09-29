package location_usecase

import (
	"time"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/prediction_usecase"
	"github.com/google/uuid"
)

type CreateLocationInputDTO struct {
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type FindLocationByIdInputDTO struct {
	Id uuid.UUID `json:"id"`
}

type CreateLocationOutputDTO struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
}

type FindLocationOutputDTO struct {
	Id          uuid.UUID                                     `json:"id"`
	Name        string                                        `json:"name"`
	Latitude    string                                        `json:"latitude"`
	Longitude   string                                        `json:"longitude"`
	Predictions []*prediction_usecase.FindPredictionOutputDTO `json:"predictions"`
	CreatedAt   time.Time                                     `json:"created_at"`
	UpdatedAt   time.Time                                     `json:"updated_at"`
}

type UpdateLocationInputDTO struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
}

type FindAllLocationsOutputDTO []*FindLocationOutputDTO

type UpdateLocationOutputDTO struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteLocationInputDTO struct {
	Id uuid.UUID `json:"id"`
}
