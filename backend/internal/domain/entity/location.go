package entity

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidLocation  = errors.New("invalid location")
	ErrLocationNotFound = errors.New("location not found")
)

type LocationRepository interface {
	CreateLocation(ctx context.Context, input *Location) (*Location, error)
	GetAllLocations(ctx context.Context) ([]*Location, error)
	GetLocationById(ctx context.Context, locationId *uuid.UUID) (*Location, error)
	UpdateLocation(ctx context.Context, input *Location) (*Location, error)
	DeleteLocation(ctx context.Context, locationId *uuid.UUID) error
}

type Location struct {
	LocationId  uuid.UUID `json:"location_id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	CoordinateX string    `json:"coordinate_x"`
	CoordinateY string    `json:"coordinate_y"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
}

func NewLocation(name *string, latitude *string, longitude *string) (*Location, error) {
	location := &Location{
		LocationId:  uuid.New(),
		Name:        *name,
		CoordinateX: *latitude,
		CoordinateY: *longitude,
		CreatedAt:   time.Now(),
	}
	if err := location.Validate(); err != nil {
		return nil, err
	}
	return location, nil
}

func (l *Location) Validate() error {
	if l.LocationId == uuid.Nil || l.Name == "" || l.CoordinateX == "" || l.CoordinateY == "" || l.CreatedAt.IsZero() {
		return ErrInvalidLocation
	}
	return nil
}
