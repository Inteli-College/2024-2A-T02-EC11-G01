package location_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
)

type FindLocationByIdUsecase struct {
	LocationRepository entity.LocationRepository
}

func NewFindLocationByIdUseCase(locationRepository entity.LocationRepository) *FindLocationByIdUsecase {
	return &FindLocationByIdUsecase{
		LocationRepository: locationRepository,
	}
}

func (u *FindLocationByIdUsecase) Execute(ctx context.Context, input *dto.FindLocationByIdInputDTO) (*dto.LocationOutputDTO, error) {
	locationUUID, errUUID := uuid.Parse(input.LocationId)
	if errUUID != nil {
		return nil, errUUID
	}

	location, err := u.LocationRepository.GetLocationById(ctx, &locationUUID)
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
