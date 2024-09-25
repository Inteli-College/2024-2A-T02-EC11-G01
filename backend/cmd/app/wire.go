//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/configs"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/event"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/event/handler"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/rabbitmq"
	web_handler "github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/web/handler"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/prediction_usecase"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/google/wire"
)

var setDBprovider = wire.NewSet(configs.SetupPostgres)

var setRabbitProvider = wire.NewSet(configs.SetupRabbitMQChannel)

var setEventDispatcher = wire.NewSet(events.NewEventDispatcher, wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewLocationCreated,
	event.NewPredictionCreated,
	wire.Bind(new(events.EventInterface), new(*event.LocationCreated)),
	wire.Bind(new(events.EventInterface), new(*event.PredictionCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setLocationRepositoryDependency = wire.NewSet(
	setDBprovider,
	repository.NewLocationRepositoryGorm,
	wire.Bind(new(entity.LocationRepository), new(*repository.LocationRepositoryGorm)),
)

var setPredictionRepositoryDependency = wire.NewSet(
	setDBprovider,
	repository.NewPredictionRepositoryGorm,
	wire.Bind(new(entity.PredictionRepository), new(*repository.PredictionRepositoryGorm)),
)

var setLocationWebHandlers = wire.NewSet(
	web_handler.NewLocationHandler,
)

var setPredictionWebHandlers = wire.NewSet(
	web_handler.NewPredictionHandler,
)

var setLocationCreatedEvent = wire.NewSet(
	event.NewLocationCreated,
	wire.Bind(new(events.EventInterface), new(*event.LocationCreated)),
)

var setPredictionCreatedEvent = wire.NewSet(
	event.NewPredictionCreated,
	wire.Bind(new(events.EventInterface), new(*event.PredictionCreated)),
)

func NewEventDispatcher() (*events.EventDispatcher, error) {
	wire.Build(setEventDispatcher)

	return nil, nil
}

func NewLocationCreatedHandler() (*handler.LocationCreatedHandler, error) {
	wire.Build(setRabbitProvider, handler.NewLocationCreatedHandler)

	return nil, nil
}

func NewPredictionCreatedHandler() (*handler.PredictionCreatedHandler, error) {
	wire.Build(setRabbitProvider, handler.NewPredictionCreatedHandler)

	return nil, nil
}

func NewRabbitMQConsumer() (*rabbitmq.RabbitMQConsumer, error) {
	wire.Build(setRabbitProvider, rabbitmq.NewRabbitMQConsumer)

	return nil, nil
}

func NewCreatePredictionUseCase() (*prediction_usecase.CreatePredictionUseCase, error) {
	wire.Build(
		setPredictionRepositoryDependency,
		setPredictionCreatedEvent,
		setEventDispatcher,
		prediction_usecase.NewCreatePredictionUseCase,
	)
	return &prediction_usecase.CreatePredictionUseCase{}, nil
}

func NewPredicitonWebHandlers() (*PredictionWebHandlers, error) {
	wire.Build(
		setPredictionRepositoryDependency,
		setPredictionCreatedEvent,
		setEventDispatcher,
		setPredictionWebHandlers,
		wire.Struct(new(PredictionWebHandlers), "*"),
	)
	return nil, nil
}

func NewLocationWebHandlers() (*LocationWebHandlers, error) {
	wire.Build(
		setLocationRepositoryDependency,
		setLocationCreatedEvent,
		setEventDispatcher,
		setLocationWebHandlers,
		wire.Struct(new(LocationWebHandlers), "*"),
	)
	return nil, nil
}

type LocationWebHandlers struct {
	LocationWebHandlers *web_handler.LocationHandler
}

type PredictionWebHandlers struct {
	PredictionWebHandlers *web_handler.PredictionHandler
}
