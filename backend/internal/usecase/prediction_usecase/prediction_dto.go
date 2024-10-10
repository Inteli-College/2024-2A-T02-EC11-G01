package prediction_usecase

import (
	"time"

	"github.com/google/uuid"
)

type CreatePredictionInputDTO struct {
	RawImagePath       string    `json:"raw_image_path"`
	AnnotatedImagePath string    `json:"annotated_image_path"`
	Detections         uint      `json:"detections"`
	LocationId         uuid.UUID `json:"location_id"`
}

type FindPredictionByIdInputDTO struct {
	PredictionId uuid.UUID `json:"prediction_id"`
}

type FindAllPredictionsByLocationIdInputDTO struct {
	LocationId uuid.UUID `json:"location_id"`
}

type UpdatePredictionInputDTO struct {
	PredictionId       uuid.UUID `json:"-" binding:"-"`
	RawImagePath       string    `json:"raw_image_path"`
	AnnotatedImagePath string    `json:"annotated_image_path"`
	Detections         uint      `json:"detections"`
	LocationId         uuid.UUID `json:"location_id"`
}

type DeletePredictionInputDTO struct {
	PredictionId uuid.UUID `json:"prediction_id"`
}

type CreatePredictionOutputDTO struct {
	PredictionId       uuid.UUID `json:"prediction_id"`
	RawImagePath       string    `json:"raw_image_path"`
	AnnotatedImagePath string    `json:"annotated_image_path"`
	Detections         uint      `json:"detections"`
	LocationId         uuid.UUID `json:"location_id"`
	CreatedAt          time.Time `json:"created_at"`
}

type FindPredictionOutputDTO struct {
	PredictionId       uuid.UUID `json:"prediction_id"`
	RawImagePath       string    `json:"raw_image_path"`
	AnnotatedImagePath string    `json:"annotated_image_path"`
	Detections         uint      `json:"detections"`
	LocationId         uuid.UUID `json:"location_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type FindAllPredictionsOutputDTO []*FindPredictionOutputDTO

type FindAllPredictionsByLocationIdOutputDTO []*FindPredictionOutputDTO

type UpdatePredictionOutputDTO struct {
	PredictionId       uuid.UUID `json:"prediction_id"`
	RawImagePath       string    `json:"raw_image_path"`
	AnnotatedImagePath string    `json:"annotated_image_path"`
	Detections         uint      `json:"detections"`
	LocationId         uuid.UUID `json:"location_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"update_at"`
}
