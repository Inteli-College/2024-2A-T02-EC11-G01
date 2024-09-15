//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/prediction_usecase"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/event"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/database"
	web_handler "github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/web/handler"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewLocationCreated,
	event.NewPredictionCreated,
	wire.Bind(new(events.EventInterface), new(*event.LocationCreated)),
	wire.Bind(new(events.EventInterface), new(*event.PredictionCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setLocationRepositoryDependency = wire.NewSet(
	database.NewLocationRepositoryGorm,
	wire.Bind(new(entity.LocationRepository), new(*database.LocationRepositoryGorm)),
)

var setPredictionRepositoryDependency = wire.NewSet(
	database.NewPredictionRepositoryGorm,
	wire.Bind(new(entity.PredictionRepository), new(*database.PredictionRepositoryGorm)),
)

var setLocationWebHandlers = wire.NewSet(
	web_handler.NewLocationHandlers,
)

var setPredictionWebHandlers = wire.NewSet(
	web_handler.NewPredictionHandlers,
)

var setLocationCreatedEvent = wire.NewSet(
	event.NewLocationCreated,
	wire.Bind(new(events.EventInterface), new(*event.LocationCreated)),
)

var setPredictionCreatedEvent = wire.NewSet(
	event.NewPredictionCreated,
	wire.Bind(new(events.EventInterface), new(*event.PredictionCreated)),
)

func NewCreatePredictionUseCase(db *gorm.DB, eventDispatcher events.EventDispatcherInterface) *prediction_usecase.CreatePredictionUseCase {
	wire.Build(
		setPredictionRepositoryDependency,
		setPredictionCreatedEvent,
		prediction_usecase.NewCreatePredictionUseCase,
	)
	return &prediction_usecase.CreatePredictionUseCase{}
}

func NewPredicitonWebHandlers(db *gorm.DB, eventDispatcher events.EventDispatcherInterface) (*PredictionWebHandlers, error) {
	wire.Build(
		setPredictionRepositoryDependency,
		setPredictionCreatedEvent,
		setPredictionWebHandlers,
		wire.Struct(new(PredictionWebHandlers), "*"),
	)
	return nil, nil
}

func NewLocationWebHandlers(db *gorm.DB, eventDispatcher events.EventDispatcherInterface) (*LocationWebHandlers, error) {
	wire.Build(
		setLocationRepositoryDependency,
		setLocationCreatedEvent,
		setLocationWebHandlers,
		wire.Struct(new(LocationWebHandlers), "*"),
	)
	return nil, nil
}

type LocationWebHandlers struct {
	LocationWebHandlers *web_handler.LocationHandlers
}

type PredictionWebHandlers struct {
	PredictionWebHandlers *web_handler.PredictionHandlers
}
