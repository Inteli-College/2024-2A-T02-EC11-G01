package location_usecase

import (
	"context"
	"time"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
)

type UpdateLocationUseCase struct {
	LocationRepository entity.LocationRepository
}

func NewUpdateLocationUseCase(locationRepository entity.LocationRepository) *UpdateLocationUseCase {
	return &UpdateLocationUseCase{
		LocationRepository: locationRepository,
	}
}

func (u *UpdateLocationUseCase) Execute(ctx context.Context, input UpdateLocationInputDTO) (*UpdateLocationOutputDTO, error) {
	location, err := u.LocationRepository.UpdateLocation(ctx, &entity.Location{
		Id:        input.Id,
		Name:      input.Name,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &UpdateLocationOutputDTO{
		Id:        location.Id,
		Name:      location.Name,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		CreatedAt: location.CreatedAt,
		UpdatedAt: location.UpdatedAt,
	}, nil
}
