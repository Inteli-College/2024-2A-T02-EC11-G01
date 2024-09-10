package usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
)

type UpdateLocationUseCase struct {
	LocationRepository repository.LocationRepository
}

func NewUpdateLocationUseCase(locationRepository repository.LocationRepository) *UpdateLocationUseCase {
	return &UpdateLocationUseCase{LocationRepository: locationRepository}
}

func (uc *UpdateLocationUseCase) Execute(ctx context.Context, locationId string, input *dto.CreateLocationInputDTO) (*dto.LocationOutputDTO, error) {
	location, err := uc.LocationRepository.UpdateLocation(ctx, locationId, input)
	if err != nil {
		return nil, err
	}
	return location, nil
}
