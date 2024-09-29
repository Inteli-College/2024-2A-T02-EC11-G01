package entity

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrPredictionNotFound = errors.New("prediction not found")
	ErrInvalidPrediction  = errors.New("invalid prediction")
)

type PredictionRepository interface {
	CreatePrediction(ctx context.Context, prediction *Prediction) (*Prediction, error)
	FindAllPredictions(ctx context.Context) ([]*Prediction, error)
	FindPredictionById(ctx context.Context, id uuid.UUID) (*Prediction, error)
	FindAllPredictionsByLocationId(ctx context.Context, id uuid.UUID) ([]*Prediction, error)
	UpdatePrediction(ctx context.Context, prediction *Prediction) (*Prediction, error)
	DeletePrediction(ctx context.Context, id uuid.UUID) error
}

type Prediction struct {
	Id             uuid.UUID `json:"id,omitempty" gorm:"primarykey;type:uuid"`
	RawImage       string    `json:"raw_image,omitempty" gorm:"type:text"`
	AnnotatedImage string    `json:"annotated_image,omitempty" gorm:"type:text"`
	Detections     uint      `json:"detections,omitempty" gorm:"type:integer"`
	LocationId     uuid.UUID `json:"location_id,omitempty" gorm:"type:uuid;not null"`
	CreatedAt      time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
}

func NewPrediction(rawImage string, annotatedImage string, detections uint, locationId uuid.UUID) (*Prediction, error) {
	prediciton := &Prediction{
		Id:             uuid.New(),
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
	if p.Id == uuid.Nil || p.RawImage == "" || p.AnnotatedImage == "" || p.LocationId == uuid.Nil || p.CreatedAt.IsZero() {
		return ErrInvalidPrediction
	}
	return nil
}
