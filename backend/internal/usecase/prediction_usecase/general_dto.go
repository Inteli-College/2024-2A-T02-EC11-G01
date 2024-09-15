package prediction_usecase

import (
	"time"

	"github.com/google/uuid"
)

type FindPredictionOutputDTO struct {
	Id             uuid.UUID `json:"id"`
	RawImage       string    `json:"raw_image"`
	AnnotatedImage string    `json:"annotated_image"`
	Detections     uint      `json:"detections"`
	LocationId     uuid.UUID `json:"location_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
