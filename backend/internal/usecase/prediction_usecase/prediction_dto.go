package prediction_usecase

import (
	"time"

	"github.com/google/uuid"
)

type CreatePredictionInputDTO struct {
	RawImage       string    `json:"raw_image"`
	AnnotatedImage string    `json:"annotated_image"`
	Detections     uint      `json:"detections"`
	LocationId     uuid.UUID `json:"location_id"`
}

type FindPredictionByIdInputDTO struct {
	Id uuid.UUID `json:"id"`
}

type FindAllPredictionsByLocationIdInputDTO struct {
	LocationId uuid.UUID `json:"location_id"`
}

type UpdatePredictionInputDTO struct {
	Id             uuid.UUID `json:"id"`
	RawImage       string    `json:"raw_image"`
	AnnotatedImage string    `json:"annotated_image"`
	Detections     uint      `json:"detections"`
	LocationId     uuid.UUID `json:"location_id"`
}

type DeletePredictionInputDTO struct {
	Id uuid.UUID `json:"id"`
}


type CreatePredictionOutputDTO struct {
	Id             uuid.UUID `json:"id"`
	RawImage       string    `json:"raw_image"`
	AnnotatedImage string    `json:"annotated_image"`
	Detections     uint      `json:"detections"`
	LocationId     uuid.UUID `json:"location_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type FindPredictionOutputDTO struct {
	Id             uuid.UUID `json:"id"`
	RawImage       string    `json:"raw_image"`
	AnnotatedImage string    `json:"annotated_image"`
	Detections     uint      `json:"detections"`
	LocationId     uuid.UUID `json:"location_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type FindAllPredictionsOutputDTO []*FindPredictionOutputDTO

type FindAllPredictionsByLocationIdOutputDTO []*FindPredictionOutputDTO

type UpdatePredictionOutputDTO struct {
	Id             uuid.UUID `json:"id"`
	RawImage       string    `json:"raw_image"`
	AnnotatedImage string    `json:"annotated_image"`
	Detections     uint      `json:"detections"`
	LocationId     uuid.UUID `json:"location_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"update_at"`
}

// type PredictionDTO struct {
// 	PredictionId    string                 `json:"prediction_id"`
// 	RawImagePath    string                 `json:"raw_image_path"`
// 	OutputImagePath string                 `json:"output_image_path"`
// 	Output          map[string]interface{} `json:"output"`
// 	LocationId      string                 `json:"location_id"`
// 	CreatedAt       time.Time              `json:"created_at"`
// }

// type CreatePredictionInputDTO struct {
// 	RawImagePath    string                `json:"raw_image_path"`
// 	OutputImagePath string                `json:"output_image_path"`
// 	Output          map[string]interface{} `json:"output"`
// 	LocationId      string                `json:"location_id"`
// }

// type FindAllPredictionsOutputDTO []*PredictionDTO

// type FindPredictionByIdInputDTO struct {
// 	PredictionId string `json:"prediction_id"`
// }

// type FindPredictionByLocationIdInputDTO struct {
// 	LocationId string `json:"location_id"`
// }
