package usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
)

type CreateLocationUseCase struct {
	LocationRepository repository.LocationRepository
}

func NewCreateLocationUseCase(locationRepository repository.LocationRepository) *CreateLocationUseCase {
	return &CreateLocationUseCase{LocationRepository: locationRepository}
}

func (uc *CreateLocationUseCase) Execute(ctx context.Context, input *dto.CreateLocationInputDTO) (*dto.LocationOutputDTO, error) {
	location, err := uc.LocationRepository.CreateLocation(ctx, input)
	if err != nil {
		return nil, err
	}
	return location, nil
}
