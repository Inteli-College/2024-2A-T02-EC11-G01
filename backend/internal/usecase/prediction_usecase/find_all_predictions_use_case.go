package prediction_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
)

type FindAllPredictionsUsecase struct {
	PredictionRepository *repository.PredictionRepository
}

func NewFindAllPredictionsUsecase(predictionRepository *repository.PredictionRepository) *FindAllPredictionsUsecase {
	return &FindAllPredictionsUsecase{PredictionRepository: predictionRepository}
}

func (uc *FindAllPredictionsUsecase) Execute(ctx context.Context, limit *int, offset *int) ([]*dto.PredictionDTO, error) {
	predictions, err := uc.PredictionRepository.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return predictions, nil
}
