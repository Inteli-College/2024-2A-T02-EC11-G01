package location_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
)

type UpdateLocationUseCase struct {
	LocationRepository entity.LocationRepository
}

func NewUpdateLocationUseCase(locationRepository entity.LocationRepository) *UpdateLocationUseCase {
	return &UpdateLocationUseCase{
		LocationRepository: locationRepository,
	}
}

func (u *UpdateLocationUseCase) Execute(ctx context.Context, input *dto.UpdateLocationInputDTO) (*dto.LocationOutputDTO, error) {
	locationUUID, errUUID := uuid.Parse(input.LocationId)
	if errUUID != nil {
		return nil, errUUID
	}

	location, errGet := u.LocationRepository.GetLocationById(ctx, &locationUUID)
	if errGet != nil {
		return nil, errGet
	}

	if input.Name != "" {
		location.Name = input.Name
	}

	if input.CoordinateX != "" {
		location.CoordinateX = input.CoordinateX
	}

	if input.CoordinateY != "" {
		location.CoordinateY = input.CoordinateY
	}

	location, err := u.LocationRepository.UpdateLocation(ctx, location)
	if err != nil {
		return nil, err
	}
	return &dto.LocationOutputDTO{
		LocationId:  location.LocationId.String(),
		Name:        location.Name,
		CoordinateX: location.CoordinateX,
		CoordinateY: location.CoordinateY,
		CreatedAt:   location.CreatedAt,
	}, nil
}
