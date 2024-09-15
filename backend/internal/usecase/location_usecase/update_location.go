package location_usecase

import (
	"time"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
)

type UpdateLocationInputDTO struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
}

type UpdateLocationOutputDTO struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateLocationUseCase struct {
	LocationRepository entity.LocationRepository
}

func NewUpdateLocationUseCase(locationRepository entity.LocationRepository) *UpdateLocationUseCase {
	return &UpdateLocationUseCase{
		LocationRepository: locationRepository,
	}
}

func (u *UpdateLocationUseCase) Execute(input UpdateLocationInputDTO) (*UpdateLocationOutputDTO, error) {
	location, err := u.LocationRepository.UpdateLocation(&entity.Location{
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
