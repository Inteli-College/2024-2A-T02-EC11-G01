package location_usecase

import (
	"context"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
)

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

func (u *CreateLocationUseCase) Execute(ctx context.Context, input *dto.CreateLocationInputDTO) (*dto.LocationOutputDTO, error) {
	location, err := entity.NewLocation(&input.Name, &input.CoordinateX, &input.CoordinateY)
	if err != nil {
		return nil, err
	}
	res, err := u.LocationRepository.CreateLocation(ctx, location)
	if err != nil {
		return nil, err
	}

	dto := &dto.LocationOutputDTO{
		LocationId:  res.LocationId.String(),
		Name:        res.Name,
		CoordinateX: res.CoordinateX,
		CoordinateY: res.CoordinateY,
		CreatedAt:   res.CreatedAt,
	}

	u.LocationCreated.SetPayload(dto)
	u.EventDispatcher.Dispatch(u.LocationCreated)

	return dto, nil
}
