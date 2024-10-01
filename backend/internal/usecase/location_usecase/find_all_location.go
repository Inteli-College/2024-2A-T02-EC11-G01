package location_usecase

import (
	"context"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/prediction_usecase"
)

type FindAllLocationsUseCase struct {
	LocationRepository entity.LocationRepository
}

func NewFindAllLocationsUseCase(locationRepository entity.LocationRepository) *FindAllLocationsUseCase {
	return &FindAllLocationsUseCase{
		LocationRepository: locationRepository,
	}
}

func (u *FindAllLocationsUseCase) Execute(ctx context.Context) (*FindAllLocationsOutputDTO, error) {
	res, err := u.LocationRepository.FindAllLocations(ctx)
	if err != nil {
		return nil, err
	}
	output := make(FindAllLocationsOutputDTO, len(res))
	for i, location := range res {
		predictions := make([]*prediction_usecase.FindPredictionOutputDTO, len(location.Predictions))
		for j, prediction := range location.Predictions {
			predictions[j] = &prediction_usecase.FindPredictionOutputDTO{
				Id:             prediction.Id,
				RawImage:       prediction.RawImage,
				AnnotatedImage: prediction.AnnotatedImage,
				Detections:     prediction.Detections,
				LocationId:     prediction.LocationId,
				CreatedAt:      prediction.CreatedAt,
				UpdatedAt:      prediction.UpdatedAt,
			}
		}
		output[i] = &FindLocationOutputDTO{
			Id:          location.Id,
			Name:        location.Name,
			Latitude:    location.Latitude,
			Longitude:   location.Longitude,
			Predictions: predictions,
			CreatedAt:   location.CreatedAt,
			UpdatedAt:   location.UpdatedAt,
		}
	}
	return &output, nil
}
