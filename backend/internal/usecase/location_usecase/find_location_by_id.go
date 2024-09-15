package location_usecase

import (
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
)

type FindLocationByIdInputDTO struct {
	Id uuid.UUID `json:"id"`
}

type FindLocationByIdUsecase struct {
	LocationRepository entity.LocationRepository
}

func NewFindLocationByIdUseCase(locationRepository entity.LocationRepository) *FindLocationByIdUsecase {
	return &FindLocationByIdUsecase{
		LocationRepository: locationRepository,
	}
}

func (u *FindLocationByIdUsecase) Execute(input FindLocationByIdInputDTO) (*FindLocationOutputDTO, error) {
	location, err := u.LocationRepository.FindLocationById(input.Id)
	if err != nil {
		return nil, err
	}
	return &FindLocationOutputDTO{
		Id:          location.Id,
		Name:        location.Name,
		CoordinateX: location.CoordinateX,
		CoordinateY: location.CoordinateY,
		CreatedAt:   location.CreatedAt,
		UpdatedAt:   location.UpdatedAt,
	}, nil
}
