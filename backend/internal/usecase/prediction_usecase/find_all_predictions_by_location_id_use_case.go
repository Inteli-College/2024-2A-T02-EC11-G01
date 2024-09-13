package prediction_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
)

type FindAllPredictionsByLocationIdUsecase struct {
	PredictionRepository *repository.PredictionRepository
}

func NewFindAllPredictionsByLocationIdUsecase(predictionRepository *repository.PredictionRepository) *FindAllPredictionsByLocationIdUsecase {
	return &FindAllPredictionsByLocationIdUsecase{PredictionRepository: predictionRepository}
}

func (uc *FindAllPredictionsByLocationIdUsecase) Execute(ctx context.Context, locationId *string, limit *int, offset *int) ([]*dto.PredictionDTO, error) {
	predictions, err := uc.PredictionRepository.GetAllByLocationID(ctx, locationId, limit, offset)
	if err != nil {
		return nil, err
	}
	return predictions, nil
}
