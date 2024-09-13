package mappers

import (
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
)

func MapPredictionEntityToDTO(prediction *entity.Prediction) *dto.PredictionDTO {
	return &dto.PredictionDTO{
		PredictionId:    prediction.PredictionId,
		OutputImagePath: prediction.OutputImagePath,
		Output:          prediction.Output,
		LocationId:      prediction.LocationId,
		RawImagePath:    prediction.RawImagePath,
		CreatedAt:       prediction.CreatedAt,
	}
}
