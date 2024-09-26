package prediction_usecase

import (
	"context"
	"encoding/json"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
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

func (u *FindAllPredictionsUseCase) Execute(ctx context.Context, limit *int, offset *int) (dto.FindAllPredictionsOutputDTO, error) {
	res, err := u.PredictionRepository.FindAllPredictions(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	output := make(dto.FindAllPredictionsOutputDTO, len(res))
	for i, prediction := range res {
		var detections map[string]interface{}
		jsonErr := json.Unmarshal(prediction.Output, &detections)
		if jsonErr != nil {
			return nil, jsonErr
		}
		output[i] = &dto.PredictionDTO{
			PredictionId:    prediction.PredictionId.String(),
			RawImagePath:    prediction.RawImagePath,
			OutputImagePath: prediction.OutputImagePath,
			Output:          detections,
			LocationId:      prediction.LocationId.String(),
			CreatedAt:       prediction.CreatedAt,
		}
	}
	return output, nil
}
