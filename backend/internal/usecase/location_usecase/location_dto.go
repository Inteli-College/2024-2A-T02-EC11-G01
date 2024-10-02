package location_usecase

import (
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/prediction_usecase"
	"github.com/google/uuid"
	"time"
)

type CreateLocationInputDTO struct {
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type FindLocationByIdInputDTO struct {
	LocationId uuid.UUID `json:"location_id"`
}

type CreateLocationOutputDTO struct {
	LocationId uuid.UUID `json:"location_id"`
	Name       string    `json:"name"`
	Latitude   string    `json:"latitude"`
	Longitude  string    `json:"longitude"`
	CreatedAt  time.Time `json:"created_at"`
}

type FindLocationOutputDTO struct {
	LocationId  uuid.UUID                                     `json:"location_id"`
	Name        string                                        `json:"name"`
	Latitude    string                                        `json:"latitude"`
	Longitude   string                                        `json:"longitude"`
	Predictions []*prediction_usecase.FindPredictionOutputDTO `json:"predictions"`
	CreatedAt   time.Time                                     `json:"created_at"`
	UpdatedAt   time.Time                                     `json:"updated_at"`
}

type UpdateLocationInputDTO struct {
	LocationId uuid.UUID `json:"location_id"`
	Name       string    `json:"name"`
	Latitude   string    `json:"latitude"`
	Longitude  string    `json:"longitude"`
}

type FindAllLocationsOutputDTO []*FindLocationOutputDTO

type UpdateLocationOutputDTO struct {
	LocationId uuid.UUID `json:"location_id"`
	Name       string    `json:"name"`
	Latitude   string    `json:"latitude"`
	Longitude  string    `json:"longitude"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type DeleteLocationInputDTO struct {
	LocationId uuid.UUID `json:"location_id"`
}
