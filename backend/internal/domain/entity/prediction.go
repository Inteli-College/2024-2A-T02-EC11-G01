package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrPredictionNotFound = errors.New("prediction not found")
	ErrInvalidPrediction  = errors.New("invalid prediction")
)

type PredictionRepository interface {
	CreatePrediction(prediction *Prediction) (*Prediction, error)
	FindAllPredictions() ([]*Prediction, error)
	FindPredictionById(id uuid.UUID) (*Prediction, error)
	UpdatePrediction(prediction *Prediction) (*Prediction, error)
	DeletePrediction(id uuid.UUID) error
}

type Prediction struct {
	PredictionId    uuid.UUID `json:"prediction_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	RawImagePath    string    `json:"raw_image_path,omitempty" gorm:"type:text;column:raw_image_path"`
	OutputImagePath string    `json:"output_image_path,omitempty" gorm:"type:text"`
	Output          []byte    `json:"Output,omitempty" gorm:"type:json"`
	LocationId      uuid.UUID `json:"location_id,omitempty" gorm:"foreignkey:LocationId;constraint:OnDelete:CASCADE;type:uuid"`
	CreatedAt       time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
}

func NewPrediction(rawImagePath string, outputImagePath string, output []byte, locationId uuid.UUID) (*Prediction, error) {
	prediciton := &Prediction{
		RawImagePath:    rawImagePath,
		OutputImagePath: outputImagePath,
		Output:          output,
		LocationId:      locationId,
		CreatedAt:       time.Now(),
	}
	if err := prediciton.Validate(); err != nil {
		return nil, err
	}
	return prediciton, nil
}

func (p *Prediction) Validate() error {
	if p.RawImagePath == "" || p.LocationId == uuid.Nil || p.CreatedAt.IsZero() {
		return ErrInvalidPrediction
	}
	return nil
}
