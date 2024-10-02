package prediction_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
)

type FindAllPredictionsUseCase struct {
	PredictionRepository entity.PredictionRepository
}

func NewFindAllPredictionsUseCase(predictionRepository entity.PredictionRepository) *FindAllPredictionsUseCase {
	return &FindAllPredictionsUseCase{
		PredictionRepository: predictionRepository,
	}
}

func (u *FindAllPredictionsUseCase) Execute(ctx context.Context) (*FindAllPredictionsOutputDTO, error) {
	res, err := u.PredictionRepository.FindAllPredictions(ctx)
	if err != nil {
		return nil, err
	}
	output := make(FindAllPredictionsOutputDTO, len(res))
	for i, prediction := range res {
		output[i] = &FindPredictionOutputDTO{
			PredictionId:       prediction.PredictionId,
			RawImagePath:       prediction.RawImagePath,
			AnnotatedImagePath: prediction.AnnotatedImagePath,
			Detections:         prediction.Detections,
			LocationId:         prediction.LocationId,
			CreatedAt:          prediction.CreatedAt,
			UpdatedAt:          prediction.UpdatedAt,
		}
	}
	return &output, nil
}
