package location_usecase

import (
	"time"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/google/uuid"
)

type CreateLocationInputDTO struct {
	Name        string `json:"name"`
	CoordinateX string `json:"coordinate_x"`
	CoordinateY string `json:"coordinate_y"`
}

type CreateLocationOutputDTO struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	CoordinateX string    `json:"coordinate_x"`
	CoordinateY string    `json:"coordinate_y"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateLocationUseCase struct {
	LocationCreated    events.EventInterface
	LocationRepository entity.LocationRepository
	EventDispatcher    events.EventDispatcherInterface
}

func NewCreateLocationUseCase(
	locationCreated events.EventInterface,
	locationRepository entity.LocationRepository,
	eventsDispatcher events.EventDispatcherInterface,
) *CreateLocationUseCase {
	return &CreateLocationUseCase{
		LocationCreated:    locationCreated,
		EventDispatcher:    eventsDispatcher,
		LocationRepository: locationRepository,
	}
}

func (u *CreateLocationUseCase) Execute(input CreateLocationInputDTO) (*CreateLocationOutputDTO, error) {
	location, err := entity.NewLocation(input.Name, input.CoordinateX, input.CoordinateY)
	if err != nil {
		return nil, err
	}
	res, err := u.LocationRepository.CreateLocation(location)
	if err != nil {
		return nil, err
	}

	dto := &CreateLocationOutputDTO{
		Id:          res.Id,
		Name:        res.Name,
		CoordinateX: res.CoordinateX,
		CoordinateY: res.CoordinateY,
		CreatedAt:   res.CreatedAt,
	}

	u.LocationCreated.SetPayload(dto)
	u.EventDispatcher.Dispatch(u.LocationCreated)

	return dto, nil
}
