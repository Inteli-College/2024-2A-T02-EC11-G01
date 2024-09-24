package location_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
)

type FindAllLocationsUseCase struct {
	LocationRepository entity.LocationRepository
}

func NewFindAllLocationsUseCase(locationRepository entity.LocationRepository) *FindAllLocationsUseCase {
	return &FindAllLocationsUseCase{
		LocationRepository: locationRepository,
	}
}

func (u *FindAllLocationsUseCase) Execute(ctx context.Context) (dto.FindAllLocationsOutputDTO, error) {
	res, err := u.LocationRepository.GetAllLocations(ctx)
	if err != nil {
		return nil, err
	}
	output := make(dto.FindAllLocationsOutputDTO, len(res))
	for i, location := range res {
		output[i] = &dto.LocationOutputDTO{
			LocationId:  location.LocationId.String(),
			Name:        location.Name,
			CoordinateX: location.CoordinateX,
			CoordinateY: location.CoordinateY,
			CreatedAt:   location.CreatedAt,
		}
	}
	return output, nil
}
