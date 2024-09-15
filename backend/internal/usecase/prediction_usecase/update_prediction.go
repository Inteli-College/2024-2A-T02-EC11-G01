package prediction_usecase

import (
	"time"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
)

type UpdatePredictionInputDTO struct {
	Id             uuid.UUID `json:"id"`
	RawImage       string    `json:"raw_image"`
	AnnotatedImage string    `json:"annotated_image"`
	Detections     uint      `json:"detections"`
	LocationId     uuid.UUID `json:"location_id"`
}

type UpdatePredictionOutputDTO struct {
	Id             uuid.UUID `json:"id"`
	RawImage       string    `json:"raw_image"`
	AnnotatedImage string    `json:"annotated_image"`
	Detections     uint      `json:"detections"`
	LocationId     uuid.UUID `json:"location_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"update_at"`
}

type UpdatePredictionUseCase struct {
	PredictionRepository entity.PredictionRepository
}

func NewUpdatePredictionUseCase(predictionRepository entity.PredictionRepository) *UpdatePredictionUseCase {
	return &UpdatePredictionUseCase{
		PredictionRepository: predictionRepository,
	}
}

func (u *UpdatePredictionUseCase) Execute(input UpdatePredictionInputDTO) (*UpdatePredictionOutputDTO, error) {
	prediction, err := u.PredictionRepository.UpdatePrediction(&entity.Prediction{
		Id:             input.Id,
		RawImage:       input.RawImage,
		AnnotatedImage: input.AnnotatedImage,
		Detections:     input.Detections,
		LocationId:     input.LocationId,
		UpdatedAt:      time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &UpdatePredictionOutputDTO{
		Id:             prediction.Id,
		RawImage:       prediction.RawImage,
		AnnotatedImage: prediction.AnnotatedImage,
		Detections:     prediction.Detections,
		LocationId:     prediction.LocationId,
		CreatedAt:      prediction.CreatedAt,
		UpdatedAt:      prediction.UpdatedAt,
	}, nil
}
