package prediction_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
)

type FindPredictionByIdUseCase struct {
	PredictionRepository entity.PredictionRepository
}

func NewFindPredictionByIdUseCase(predictionRepository entity.PredictionRepository) *FindPredictionByIdUseCase {
	return &FindPredictionByIdUseCase{
		PredictionRepository: predictionRepository,
	}
}

func (u *FindPredictionByIdUseCase) Execute(ctx context.Context, input FindPredictionByIdInputDTO) (*FindPredictionOutputDTO, error) {
	prediction, err := u.PredictionRepository.FindPredictionById(ctx, input.PredictionId)
	if err != nil {
		return nil, err
	}
	return &FindPredictionOutputDTO{
		PredictionId:       prediction.PredictionId,
		RawImagePath:       prediction.RawImagePath,
		AnnotatedImagePath: prediction.AnnotatedImagePath,
		Detections:         prediction.Detections,
		LocationId:         prediction.LocationId,
		CreatedAt:          prediction.CreatedAt,
		UpdatedAt:          prediction.UpdatedAt,
	}, nil
}
