//go:build wireinject
// +build wireinject

package main

//
// import (
// 	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
// 	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/interfaces/http"
// 	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase"
// 	"github.com/google/wire"
// 	"gorm.io/gorm"
// )
//
// func InitializeLocationHandler(db *gorm.DB) *http.LocationHandler {
// 	wire.Build(
// 		repository.NewLocationRepository,
// 		usecase.NewCreateLocationUseCase,
// 		usecase.NewFindLocationByIdUseCase,
// 		usecase.NewUpdateLocationUseCase,
// 		usecase.NewDeleteLocationUseCase,
// 		usecase.NewFindAllLocationsUseCase,
// 		http.NewLocationHandler,
// 	)
// 	return &http.LocationHandler{}
// }
