// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/event"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/database"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/web/handler"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/prediction_usecase"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func NewCreatePredictionUseCase(db *gorm.DB, eventDispatcher events.EventDispatcherInterface) *prediction_usecase.CreatePredictionUseCase {
	predictionCreated := event.NewPredictionCreated()
	predictionRepositoryGorm := database.NewPredictionRepositoryGorm(db)
	createPredictionUseCase := prediction_usecase.NewCreatePredictionUseCase(predictionCreated, predictionRepositoryGorm, eventDispatcher)
	return createPredictionUseCase
}

func NewPredicitonWebHandlers(db *gorm.DB, eventDispatcher events.EventDispatcherInterface) (*PredictionWebHandlers, error) {
	predictionRepositoryGorm := database.NewPredictionRepositoryGorm(db)
	predictionCreated := event.NewPredictionCreated()
	predictionHandlers := handler.NewPredictionHandlers(eventDispatcher, predictionRepositoryGorm, predictionCreated)
	predictionWebHandlers := &PredictionWebHandlers{
		PredictionWebHandlers: predictionHandlers,
	}
	return predictionWebHandlers, nil
}

func NewLocationWebHandlers(db *gorm.DB, eventDispatcher events.EventDispatcherInterface) (*LocationWebHandlers, error) {
	locationRepositoryGorm := database.NewLocationRepositoryGorm(db)
	locationCreated := event.NewLocationCreated()
	locationHandlers := handler.NewLocationHandlers(eventDispatcher, locationRepositoryGorm, locationCreated)
	locationWebHandlers := &LocationWebHandlers{
		LocationWebHandlers: locationHandlers,
	}
	return locationWebHandlers, nil
}

// wire.go:

var setEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewLocationCreated, event.NewPredictionCreated, wire.Bind(new(events.EventInterface), new(*event.LocationCreated)), wire.Bind(new(events.EventInterface), new(*event.PredictionCreated)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setLocationRepositoryDependency = wire.NewSet(database.NewLocationRepositoryGorm, wire.Bind(new(entity.LocationRepository), new(*database.LocationRepositoryGorm)))

var setPredictionRepositoryDependency = wire.NewSet(database.NewPredictionRepositoryGorm, wire.Bind(new(entity.PredictionRepository), new(*database.PredictionRepositoryGorm)))

var setLocationWebHandlers = wire.NewSet(handler.NewLocationHandlers)

var setPredictionWebHandlers = wire.NewSet(handler.NewPredictionHandlers)

var setLocationCreatedEvent = wire.NewSet(event.NewLocationCreated, wire.Bind(new(events.EventInterface), new(*event.LocationCreated)))

var setPredictionCreatedEvent = wire.NewSet(event.NewPredictionCreated, wire.Bind(new(events.EventInterface), new(*event.PredictionCreated)))

type LocationWebHandlers struct {
	LocationWebHandlers *handler.LocationHandlers
}

type PredictionWebHandlers struct {
	PredictionWebHandlers *handler.PredictionHandlers
}