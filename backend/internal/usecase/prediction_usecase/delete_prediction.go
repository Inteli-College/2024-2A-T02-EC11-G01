package prediction_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
)

type DeletePredictionUseCase struct {
	PredictionRepository entity.PredictionRepository
}

func NewDeletePredictionUseCase(predictionRepository entity.PredictionRepository) *DeletePredictionUseCase {
	return &DeletePredictionUseCase{
		PredictionRepository: predictionRepository,
	}
}

func (u *DeletePredictionUseCase) Execute(ctx context.Context, input DeletePredictionInputDTO) error {
	return u.PredictionRepository.DeletePrediction(ctx, input.PredictionId)
}
