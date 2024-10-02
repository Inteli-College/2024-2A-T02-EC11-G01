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
	CreateLocation(ctx context.Context, location *Location) (*Location, error)
	FindAllLocations(ctx context.Context) ([]*Location, error)
	FindLocationById(ctx context.Context, id uuid.UUID) (*Location, error)
	UpdateLocation(ctx context.Context, location *Location) (*Location, error)
	DeleteLocation(ctx context.Context, id uuid.UUID) error
}

type Location struct {
	LocationId uuid.UUID     `json:"location_id,omitempty" gorm:"primarykey;type:uuid"`
	Name        string        `json:"name" gorm:"type:text"`
	Latitude    string        `json:"latitude" gorm:"type:text"`
	Longitude   string        `json:"longitude" gorm:"type:text"`
	Predictions []*Prediction `json:"predictions,omitempty" gorm:"foreignKey:LocationId;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time     `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty" gorm:"type:timestamp"`
}

func NewLocation(name string, latitude string, longitude string) (*Location, error) {
	location := &Location{
		LocationId: uuid.New(),
		Name:        name,
		Latitude:    latitude,
		Longitude:   longitude,
		CreatedAt:   time.Now(),
	}
	if err := location.Validate(); err != nil {
		return nil, err
	}
	return location, nil
}

func (l *Location) Validate() error {
	if l.LocationId == uuid.Nil || l.Name == "" || l.Latitude == "" || l.Longitude == "" || l.CreatedAt.IsZero() {
		return ErrInvalidLocation
	}
	return nil
}
