package dto

import (
	"time"
)

type FindLocationByIdInputDTO struct {
	LocationId string `json:"location_id"`
}

type CreateLocationInputDTO struct {
	Name        string `json:"name"`
	CoordinateX string `json:"coordinate_x"`
	CoordinateY string `json:"coordinate_y"`
}

type LocationOutputDTO struct {
	LocationId  string    `json:"location_id"`
	Name        string    `json:"name"`
	CoordinateX string    `json:"coordinate_x"`
	CoordinateY string    `json:"coordinate_y"`
	CreatedAt   time.Time `json:"created_at"`
}

type DeleteLocationInputDTO struct {
	LocationId string `json:"location_id"`
}

type FindAllLocationsOutputDTO []*LocationOutputDTO

type UpdateLocationInputDTO struct {
	LocationId string `json:"location_id"`
	Name       string `json:"name"`
	CoordinateX   string `json:"coordinate_x"`
	CoordinateY  string `json:"coordinate_y"`
}
