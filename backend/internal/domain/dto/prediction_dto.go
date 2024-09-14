package dto

import "time"

type PredictionDTO struct {
	PredictionId    string                 `json:"prediction_id"`
	RawImagePath    string                 `json:"raw_image_path"`
	OutputImagePath string                 `json:"output_image_path"`
	Output          map[string]interface{} `json:"output"`
	LocationId      string                 `json:"location_id"`
	CreatedAt       time.Time              `json:"created_at"`
}

type CreatePredictionInputDTO struct {
	RawImagePath    *string                `json:"raw_image_path"`
	OutputImagePath *string                `json:"output_image_path"`
	Output          map[string]interface{} `json:"output"`
	LocationId      *string                `json:"location_id"`
}
