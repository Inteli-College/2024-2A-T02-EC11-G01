package prediction_usecase

import (
	"time"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/google/uuid"
)

type CreatePredictionInputDTO struct {
	RawImage       string    `json:"raw_image"`
	AnnotatedImage string    `json:"annotated_image"`
	Detections     uint      `json:"detections"`
	LocationId     uuid.UUID `json:"location_id"`
}

type CreatePredictionOutputDTO struct {
	Id             uuid.UUID `json:"id"`
	RawImage       string    `json:"raw_image"`
	AnnotatedImage string    `json:"annotated_image"`
	Detections     uint      `json:"detections"`
	LocationId     uuid.UUID `json:"location_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreatePredictionUseCase struct {
	PredictionCreated events.EventInterface
	PredictionRepository entity.PredictionRepository
	EventDispatcher     events.EventDispatcherInterface
}

func NewCreatePredictionUseCase(
	predictionCreated events.EventInterface,
	predictionRepository entity.PredictionRepository,
	eventsDispatcher events.EventDispatcherInterface,
) *CreatePredictionUseCase {
	return &CreatePredictionUseCase{
		PredictionCreated:    predictionCreated,
		EventDispatcher:      eventsDispatcher,
		PredictionRepository: predictionRepository,
	}
}

func (u *CreatePredictionUseCase) Execute(input CreatePredictionInputDTO) (*CreatePredictionOutputDTO, error) {
	prediction, err := entity.NewPrediction(input.RawImage, input.AnnotatedImage, input.Detections, input.LocationId)
	if err != nil {
		return nil, err
	}
	res, err := u.PredictionRepository.CreatePrediction(prediction)
	if err != nil {
		return nil, err
	}
	dto := &CreatePredictionOutputDTO{
		Id:             res.Id,
		RawImage:       res.RawImage,
		AnnotatedImage: res.AnnotatedImage,
		Detections:     res.Detections,
		LocationId:     res.LocationId,
		CreatedAt:      res.CreatedAt,
	}

	u.PredictionCreated.SetPayload(dto)
	u.EventDispatcher.Dispatch(u.PredictionCreated)

	return dto, nil
}
