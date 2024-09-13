package prediction_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
)

type FindPredictionByIdUseCase struct {
	PredictionRepository repository.PredictionRepository
}

func NewFindPredictionByIdUsecase(predictionRepository repository.PredictionRepository) *FindPredictionByIdUseCase {
	return &FindPredictionByIdUseCase{PredictionRepository: predictionRepository}
}

func (uc *FindPredictionByIdUseCase) Execute(ctx context.Context, predictionId *string) (*dto.PredictionDTO, error) {
	prediction, err := uc.PredictionRepository.GetByID(ctx, predictionId)
	if err != nil {
		return nil, err
	}
	return prediction, nil
}
