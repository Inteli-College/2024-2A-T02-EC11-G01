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
	Id             uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	RawImage       string    `json:"raw_image,omitempty" gorm:"type:text"`
	AnnotatedImage string    `json:"annotated_image,omitempty" gorm:"type:text"`
	Detections     uint      `json:"detections,omitempty" gorm:"type:integer"`
	LocationId     uuid.UUID `json:"location_id,omitempty" gorm:"foreignkey:LocationId;constraint:OnDelete:CASCADE;type:uuid"`
	CreatedAt      time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
}

func NewPrediction(rawImage string, annotatedImage string, detections uint, locationId uuid.UUID) (*Prediction, error) {
	prediciton := &Prediction{
		RawImage:       rawImage,
		AnnotatedImage: annotatedImage,
		Detections:     detections,
		LocationId:     locationId,
		CreatedAt:      time.Now(),
	}
	if err := prediciton.Validate(); err != nil {
		return nil, err
	}
	return prediciton, nil
}

func (p *Prediction) Validate() error {
	if p.RawImage == "" || p.AnnotatedImage == "" || p.LocationId == uuid.Nil || p.CreatedAt.IsZero() {
		return ErrInvalidPrediction
	}
	return nil
}
