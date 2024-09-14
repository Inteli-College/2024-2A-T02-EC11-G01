package mappers

import (
	"encoding/json"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
)

func MapPredictionEntityToDTO(prediction *entity.Prediction) *dto.PredictionDTO {
	var predictionOutput map[string]interface{}

	// TODO: handle this error
	json.Unmarshal(prediction.Output, &predictionOutput)

	return &dto.PredictionDTO{
		PredictionId:    prediction.PredictionId,
		OutputImagePath: prediction.OutputImagePath,
		Output:          predictionOutput,
		LocationId:      prediction.LocationId,
		RawImagePath:    prediction.RawImagePath,
		CreatedAt:       prediction.CreatedAt,
	}
}
