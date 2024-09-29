package prediction_usecase

import (
	"context"
	"time"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
)

type UpdatePredictionUseCase struct {
	PredictionRepository entity.PredictionRepository
}

func NewUpdatePredictionUseCase(predictionRepository entity.PredictionRepository) *UpdatePredictionUseCase {
	return &UpdatePredictionUseCase{
		PredictionRepository: predictionRepository,
	}
}

func (u *UpdatePredictionUseCase) Execute(ctx context.Context, input UpdatePredictionInputDTO) (*UpdatePredictionOutputDTO, error) {
	prediction, err := u.PredictionRepository.UpdatePrediction(ctx, &entity.Prediction{
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
