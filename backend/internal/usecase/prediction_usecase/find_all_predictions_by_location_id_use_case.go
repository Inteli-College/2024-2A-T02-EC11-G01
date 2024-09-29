package prediction_usecase

import (
	"context"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
)

type FindAllPredictionsByLocationIdUseCase struct {
	PredictionRepository entity.PredictionRepository
}

func NewFindAllPredictionsByLocationIdUseCase(predictionRepository entity.PredictionRepository) *FindAllPredictionsByLocationIdUseCase {
	return &FindAllPredictionsByLocationIdUseCase{PredictionRepository: predictionRepository}
}

func (uc *FindAllPredictionsByLocationIdUseCase) Execute(ctx context.Context, input FindAllPredictionsByLocationIdInputDTO) (*FindAllPredictionsByLocationIdOutputDTO, error) {
	res, err := uc.PredictionRepository.FindAllPredictionsByLocationId(ctx, input.LocationId)
	if err != nil {
		return nil, err
	}
	output := make(FindAllPredictionsByLocationIdOutputDTO, len(res))
	for i, prediction := range res {
		output[i] = &FindPredictionOutputDTO{
			Id:             prediction.Id,
			RawImage:       prediction.RawImage,
			AnnotatedImage: prediction.AnnotatedImage,
			Detections:     prediction.Detections,
			LocationId:     prediction.LocationId,
			CreatedAt:      prediction.CreatedAt,
			UpdatedAt:      prediction.UpdatedAt,
		}
	}
	return &output, nil
}
