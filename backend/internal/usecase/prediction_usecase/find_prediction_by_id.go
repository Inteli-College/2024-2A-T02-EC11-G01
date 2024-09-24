package prediction_usecase

import (
	"context"
	"encoding/json"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
)

type FindPredictionByIdUseCase struct {
	PredictionRepository entity.PredictionRepository
}

func NewFindPredictionByIdUseCase(predictionRepository entity.PredictionRepository) *FindPredictionByIdUseCase {
	return &FindPredictionByIdUseCase{
		PredictionRepository: predictionRepository,
	}
}

func (u *FindPredictionByIdUseCase) Execute(ctx context.Context, input *dto.FindPredictionByIdInputDTO) (*dto.PredictionDTO, error) {
	predictionUUID, errUUID := uuid.Parse(input.PredictionId)
	if errUUID != nil {
		return nil, errUUID
	}

	prediction, err := u.PredictionRepository.FindPredictionById(ctx, &predictionUUID)
	if err != nil {
		return nil, err
	}

	var outputJson map[string]interface{}

	unmarshalErr := json.Unmarshal(prediction.Output, &outputJson)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &dto.PredictionDTO{
		PredictionId:    prediction.PredictionId.String(),
		RawImagePath:    prediction.RawImagePath,
		OutputImagePath: prediction.OutputImagePath,
		Output:          outputJson,
		LocationId:      prediction.LocationId.String(),
		CreatedAt:       prediction.CreatedAt,
	}, nil
}
