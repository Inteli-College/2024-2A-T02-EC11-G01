package prediction_usecase

import (
	"context"
	"encoding/json"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
	"github.com/google/uuid"
)

type FindAllPredictionsByLocationIdUsecase struct {
	PredictionRepository *repository.PredictionRepositoryGorm
}

func NewFindAllPredictionsByLocationIdUsecase(predictionRepository *repository.PredictionRepositoryGorm) *FindAllPredictionsByLocationIdUsecase {
	return &FindAllPredictionsByLocationIdUsecase{PredictionRepository: predictionRepository}
}

func (uc *FindAllPredictionsByLocationIdUsecase) Execute(ctx context.Context, input *dto.FindPredictionByLocationIdInputDTO, limit *int, offset *int) ([]*dto.PredictionDTO, error) {
	locationUUID, errUUID := uuid.Parse(input.LocationId)
	if errUUID != nil {
		return nil, errUUID
	}

	predictions, err := uc.PredictionRepository.FindAllPredictionsByLocationID(ctx, &locationUUID, limit, offset)
	if err != nil {
		return nil, err
	}

	output := make(dto.FindAllPredictionsOutputDTO, len(predictions))
	for i, prediction := range predictions {
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
