package usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
)

type DeleteLocationUseCase struct {
	LocationRepository repository.LocationRepository
}

func NewDeleteLocationUseCase(locationRepository repository.LocationRepository) *DeleteLocationUseCase {
	return &DeleteLocationUseCase{LocationRepository: locationRepository}
}

func (uc *DeleteLocationUseCase) Execute(ctx context.Context, locationId string) error {
	err := uc.LocationRepository.DeleteLocation(ctx, locationId)
	if err != nil {
		return err
	}
	return nil
}
