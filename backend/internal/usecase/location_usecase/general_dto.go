package location_usecase

import (
	"time"

	"github.com/google/uuid"
)

type FindLocationOutputDTO struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
