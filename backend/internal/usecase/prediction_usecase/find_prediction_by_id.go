package prediction_usecase

import (
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
)

type FindPredictionByIdInputDTO struct {
	Id uuid.UUID `json:"id"`
}

type FindPredictionByIdUseCase struct {
	PredictionRepository entity.PredictionRepository
}

func NewFindPredictionByIdUseCase(predictionRepository entity.PredictionRepository) *FindPredictionByIdUseCase {
	return &FindPredictionByIdUseCase{
		PredictionRepository: predictionRepository,
	}
}

func (u *FindPredictionByIdUseCase) Execute(input FindPredictionByIdInputDTO) (*FindPredictionOutputDTO, error) {
	prediction, err := u.PredictionRepository.FindPredictionById(input.Id)
	if err != nil {
		return nil, err
	}
	return &FindPredictionOutputDTO{
		Id:             prediction.Id,
		RawImage:       prediction.RawImage,
		AnnotatedImage: prediction.AnnotatedImage,
		Detections:     prediction.Detections,
		LocationId:     prediction.LocationId,
		CreatedAt:      prediction.CreatedAt,
		UpdatedAt:      prediction.UpdatedAt,
	}, nil
}
