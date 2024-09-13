//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/configs"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/web"
	"github.com/google/wire"
)

var DBProviderSet = wire.NewSet(configs.SetupPostgres)
var LocationRepositorySet = wire.NewSet(DBProviderSet, repository.NewLocationRepository)
var PredictionRepositorySet = wire.NewSet(DBProviderSet, repository.NewPredictionRepository)

func InitializeLocationsHandler() (*web.LocationHandler, error) {
	wire.Build(LocationRepositorySet, web.NewLocationHandler)

	return nil, nil
}

func InitializePredictionsHandler() (*web.PredictionHandler, error) {
	wire.Build(PredictionRepositorySet, web.NewPredictionHandler)

	return nil, nil
}
