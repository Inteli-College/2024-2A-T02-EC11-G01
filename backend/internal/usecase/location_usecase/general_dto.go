package location_usecase

import (
	"time"

	"github.com/google/uuid"
)

type FindLocationOutputDTO struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	CoordinateX string    `json:"coordinate_x"`
	CoordinateY string    `json:"coordinate_y"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
