package location_usecase

import (
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
)

type DeleteLocationInputDTO struct {
	Id uuid.UUID `json:"id"`
}

type DeleteLocationUseCase struct {
	LocationRepository entity.LocationRepository
}

func NewDeleteLocationUseCase(locationRepository entity.LocationRepository) *DeleteLocationUseCase {
	return &DeleteLocationUseCase{
		LocationRepository: locationRepository,
	}
}

func (u *DeleteLocationUseCase) Execute(input DeleteLocationInputDTO) error {
	return u.LocationRepository.DeleteLocation(input.Id)
}
