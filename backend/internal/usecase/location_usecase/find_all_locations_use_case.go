package usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
)

type FindAllLocationsUseCase struct {
	LocationRepository *repository.LocationRepository
}

func NewFindAllLocationsUseCase(locationRepository *repository.LocationRepository) *FindAllLocationsUseCase {
	return &FindAllLocationsUseCase{LocationRepository: locationRepository}
}

func (uc *FindAllLocationsUseCase) Execute(ctx context.Context) ([]*dto.LocationOutputDTO, error) {
	locations, err := uc.LocationRepository.GetAllLocations(ctx)
	if err != nil {
		return nil, err
	}
	return locations, nil
}
