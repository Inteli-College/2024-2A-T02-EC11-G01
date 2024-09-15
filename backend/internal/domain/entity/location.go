package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidLocation  = errors.New("invalid location")
	ErrLocationNotFound = errors.New("location not found")
)

type LocationRepository interface {
	CreateLocation(location *Location) (*Location, error)
	FindAllLocations() ([]*Location, error)
	FindLocationById(id uuid.UUID) (*Location, error)
	UpdateLocation(location *Location) (*Location, error)
	DeleteLocation(id uuid.UUID) error
}

type Location struct {
	LocationId  uuid.UUID `json:"location_id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	CoordinateX string    `json:"coordinate_x"`
	CoordinateY string    `json:"coordinate_y"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
}

func NewLocation(name string, coordinateX string, coordinateY string) (*Location, error) {
	location := &Location{
		Name:        name,
		CoordinateX: coordinateX,
		CoordinateY: coordinateY,
		CreatedAt:   time.Now(),
	}
	if err := location.Validate(); err != nil {
		return nil, err
	}
	return location, nil
}

func (l *Location) Validate() error {
	if l.Name == "" || l.CoordinateX == "" || l.CoordinateY == "" || l.CreatedAt.IsZero() {
		return ErrInvalidLocation
	}
	return nil
}
