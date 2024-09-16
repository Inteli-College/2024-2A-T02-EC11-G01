package prediction_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
)

type CreatePredictionUseCase struct {
	PredictionRepository *repository.PredictionRepository
}

func NewCreatePredictionUsecase(predictionRepository *repository.PredictionRepository) *CreatePredictionUseCase {
	return &CreatePredictionUseCase{PredictionRepository: predictionRepository}
}

func (uc *CreatePredictionUseCase) Execute(ctx context.Context, input *dto.CreatePredictionInputDTO) (*dto.PredictionDTO, error) {
	prediction, err := uc.PredictionRepository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return prediction, nil
}
