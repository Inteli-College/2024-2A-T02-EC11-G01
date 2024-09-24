package prediction_usecase

import (
	"context"
	"encoding/json"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/google/uuid"
)

type CreatePredictionUseCase struct {
	PredictionCreated    events.EventInterface
	PredictionRepository entity.PredictionRepository
	EventDispatcher      events.EventDispatcherInterface
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

func (u *CreatePredictionUseCase) Execute(ctx context.Context, input *dto.CreatePredictionInputDTO) (*dto.PredictionDTO, error) {
	outputJson, jsonErr := json.Marshal(input.Output)
	if jsonErr != nil {
		return nil, jsonErr
	}

	locationUUID, UUIDErr := uuid.Parse(input.LocationId)
	if UUIDErr != nil {
		return nil, UUIDErr
	}

	prediction, err := entity.NewPrediction(&input.RawImagePath, &input.OutputImagePath, outputJson, &locationUUID)
	if err != nil {
		return nil, err
	}

	res, err := u.PredictionRepository.CreatePrediction(ctx, prediction)
	if err != nil {
		return nil, err
	}

	var output map[string]interface{}

	unmarshalErr := json.Unmarshal(res.Output, &output)
	if unmarshalErr != nil {
		return nil, jsonErr
	}

	dto := &dto.PredictionDTO{
		PredictionId:    res.PredictionId.String(),
		RawImagePath:    res.RawImagePath,
		OutputImagePath: res.OutputImagePath,
		Output:          output,
		LocationId:      res.LocationId.String(),
		CreatedAt:       res.CreatedAt,
	}

	u.PredictionCreated.SetPayload(dto)
	u.EventDispatcher.Dispatch(u.PredictionCreated)

	return dto, nil
}
