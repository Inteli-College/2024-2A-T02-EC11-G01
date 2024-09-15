package prediction_usecase

import (
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
)

type DeletePredictionInputDTO struct {
	Id uuid.UUID `json:"id"`
}

type DeletePredictionUseCase struct {
	PredictionRepository entity.PredictionRepository
}

func NewDeletePredictionUseCase(predictionRepository entity.PredictionRepository) *DeletePredictionUseCase {
	return &DeletePredictionUseCase{
		PredictionRepository: predictionRepository,
	}
}

func (u *DeletePredictionUseCase) Execute(input DeletePredictionInputDTO) error {
	return u.PredictionRepository.DeletePrediction(input.Id)
}
