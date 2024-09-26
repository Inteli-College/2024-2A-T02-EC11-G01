package location_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
)

type DeleteLocationUseCase struct {
	LocationRepository entity.LocationRepository
}

func NewDeleteLocationUseCase(locationRepository entity.LocationRepository) *DeleteLocationUseCase {
	return &DeleteLocationUseCase{
		LocationRepository: locationRepository,
	}
}

func (u *DeleteLocationUseCase) Execute(ctx context.Context, input *dto.DeleteLocationInputDTO) error {
	locationUUID, err := uuid.Parse(input.LocationId)
	if err != nil {
		return err
	}

	return u.LocationRepository.DeleteLocation(ctx, &locationUUID)
}
