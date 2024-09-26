package handler

import (
	"context"
	"net/http"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
	usecase "github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/location_usecase"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	createLocationUseCase   *usecase.CreateLocationUseCase
	findLocationByIdUseCase *usecase.FindLocationByIdUsecase
	updateLocationUseCase   *usecase.UpdateLocationUseCase
	deleteLocationUseCase   *usecase.DeleteLocationUseCase
	findAllLocationsUseCase *usecase.FindAllLocationsUseCase
}

func NewLocationHandler(locationRepo *repository.LocationRepositoryGorm, eventDispatcher events.EventDispatcherInterface,
	locationRepository entity.LocationRepository, locationCreatedEvent events.EventInterface) *LocationHandler {
	handler := &LocationHandler{
		createLocationUseCase:   usecase.NewCreateLocationUseCase(locationCreatedEvent, locationRepo, eventDispatcher),
		findLocationByIdUseCase: usecase.NewFindLocationByIdUseCase(locationRepo),
		updateLocationUseCase:   usecase.NewUpdateLocationUseCase(locationRepo),
		deleteLocationUseCase:   usecase.NewDeleteLocationUseCase(locationRepo),
		findAllLocationsUseCase: usecase.NewFindAllLocationsUseCase(locationRepo),
	}
	return handler
}

// CreateLocation godoc
// @Summary Create a new location
// @Description Create a new location
// @Tags locations
// @Accept json
// @Produce json
// @Param location body dto.CreateLocationInputDTO true "Create Location Input"
// @Success 201 {object} dto.LocationOutputDTO
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /locations [post]
func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var input dto.CreateLocationInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	location, err := h.createLocationUseCase.Execute(ctx, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, location)
}

// FindAllLocations godoc
// @Summary Get all locations
// @Description Get a list of all locations
// @Tags locations
// @Accept json
// @Produce json
// @Success 200 {array} dto.LocationOutputDTO
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /locations [get]
func (h *LocationHandler) FindAllLocations(c *gin.Context) {
	ctx := context.Background()
	locations, err := h.findAllLocationsUseCase.Execute(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, locations)
}

// FindLocationById godoc
// @Summary Get a location by ID
// @Description Get a location by ID
// @Tags locations
// @Accept json
// @Produce json
// @Param id path string true "Location ID"
// @Success 200 {object} dto.LocationOutputDTO
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /locations/{id} [get]
func (h *LocationHandler) FindLocationById(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()

	var input dto.FindLocationByIdInputDTO

	input.LocationId = id

	location, err := h.findLocationByIdUseCase.Execute(ctx, &input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, location)
}

// UpdateLocation godoc
// @Summary Update a location
// @Description Update a location by ID
// @Tags locations
// @Accept json
// @Produce json
// @Param id path string true "Location ID"
// @Param location body dto.CreateLocationInputDTO true "Update Location Input"
// @Success 200 {object} dto.LocationOutputDTO
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /locations/{id} [put]
func (h *LocationHandler) UpdateLocation(c *gin.Context) {
	id := c.Param("id")

	var input dto.UpdateLocationInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.LocationId = id

	ctx := context.Background()
	location, err := h.updateLocationUseCase.Execute(ctx, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, location)
}

// DeleteLocation godoc
// @Summary Delete a location
// @Description Delete a location by ID
// @Tags locations
// @Accept json
// @Produce json
// @Param id path string true "Location ID"
// @Success 204 {object} map[string]string "Success"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /locations/{id} [delete]
func (h *LocationHandler) DeleteLocation(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()

	var input dto.DeleteLocationInputDTO
	input.LocationId = id

	err := h.deleteLocationUseCase.Execute(ctx, &input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
