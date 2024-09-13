package usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
)

type FindLocationByIdUseCase struct {
	LocationRepository *repository.LocationRepository
}

func NewFindLocationByIdUseCase(locationRepository *repository.LocationRepository) *FindLocationByIdUseCase {
	return &FindLocationByIdUseCase{LocationRepository: locationRepository}
}

func (uc *FindLocationByIdUseCase) Execute(ctx context.Context, locationId string) (*dto.LocationOutputDTO, error) {
	location, err := uc.LocationRepository.GetLocationById(ctx, locationId)
	if err != nil {
		return nil, err
	}
	return location, nil
}
