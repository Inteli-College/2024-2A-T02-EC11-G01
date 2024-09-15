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
	Id        uuid.UUID `json:"id,omitempty" gorm:"primarykey;type:uuid"`
	Name      string    `json:"name" gorm:"type:text"`
	Latitude  string    `json:"latitude" gorm:"type:text"`
	Longitude string    `json:"longitude" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
}

func NewLocation(name string, latitude string, longitude string) (*Location, error) {
	location := &Location{
		Id:        uuid.New(),
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
		CreatedAt: time.Now(),
	}
	if err := location.Validate(); err != nil {
		return nil, err
	}
	return location, nil
}

func (l *Location) Validate() error {
	if l.Id == uuid.Nil || l.Name == "" || l.Latitude == "" || l.Longitude == "" || l.CreatedAt.IsZero() {
		return ErrInvalidLocation
	}
	return nil
}
