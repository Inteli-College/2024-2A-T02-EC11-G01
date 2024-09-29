package location_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/prediction_usecase"
)

type FindLocationByIdUsecase struct {
	LocationRepository entity.LocationRepository
}

func NewFindLocationByIdUseCase(locationRepository entity.LocationRepository) *FindLocationByIdUsecase {
	return &FindLocationByIdUsecase{
		LocationRepository: locationRepository,
	}
}

func (u *FindLocationByIdUsecase) Execute(ctx context.Context, input FindLocationByIdInputDTO) (*FindLocationOutputDTO, error) {
	location, err := u.LocationRepository.FindLocationById(ctx, input.Id)
	if err != nil {
		return nil, err
	}
	var predictions []*prediction_usecase.FindPredictionOutputDTO
	for _, prediction := range location.Predictions {
		predictions = append(predictions, &prediction_usecase.FindPredictionOutputDTO{
			Id:             prediction.Id,
			RawImage:       prediction.RawImage,
			AnnotatedImage: prediction.AnnotatedImage,
			Detections:     prediction.Detections,
			LocationId:     prediction.LocationId,
			CreatedAt:      prediction.CreatedAt,
			UpdatedAt:      prediction.UpdatedAt,
		})
	}
	return &FindLocationOutputDTO{
		Id:          location.Id,
		Name:        location.Name,
		Latitude:    location.Latitude,
		Longitude:   location.Longitude,
		Predictions: predictions,
		CreatedAt:   location.CreatedAt,
		UpdatedAt:   location.UpdatedAt,
	}, nil
}
