package handler

import (
	"context"
	"net/http"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/location_usecase"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LocationHandlers struct {
	EventDispatcher      events.EventDispatcherInterface
	LocationRepository   entity.LocationRepository
	LocationCreatedEvent events.EventInterface
}

func NewLocationHandlers(
	eventDispatcher events.EventDispatcherInterface,
	locationRepository entity.LocationRepository,
	locationCreatedEvent events.EventInterface,
) *LocationHandlers {
	return &LocationHandlers{
		EventDispatcher:      eventDispatcher,
		LocationRepository:   locationRepository,
		LocationCreatedEvent: locationCreatedEvent,
	}
}

// CreateLocationHandler
// @Summary Create a new Location
// @Description Create a new Location in the system
// @Tags Locations
// @Accept json
// @Produce json
// @Param input body location_usecase.CreateLocationInputDTO true "Location entity to create"
// @Success 200 {object} location_usecase.CreateLocationOutputDTO
// @Router /location [post]
func (h *LocationHandlers) CreateLocationHandler(c *gin.Context) {
	var input location_usecase.CreateLocationInputDTO
	ctx := context.Background()
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := location_usecase.NewCreateLocationUseCase(
		h.LocationCreatedEvent,
		h.LocationRepository,
		h.EventDispatcher,
	).Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// FindLocationByIdHandler
// @Summary Retrieve a Location by ID
// @Description Get details of a specific Location by its ID
// @Tags Locations
// @Accept json
// @Produce json
// @Param location_id path string true "Location ID"
// @Success 200 {object} location_usecase.FindLocationOutputDTO
// @Router /location/{location_id} [get]
func (h *LocationHandlers) FindLocationByIdHandler(c *gin.Context) {
	var input location_usecase.FindLocationByIdInputDTO
	locationId, err := uuid.Parse(c.Param("location_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	input.LocationId = locationId

	ctx := context.Background()
	ctx = context.WithValue(ctx, "limit", c.DefaultQuery("limit", "20"))
	ctx = context.WithValue(ctx, "offset", c.DefaultQuery("offset", "0"))
	res, err := location_usecase.NewFindLocationByIdUseCase(h.LocationRepository).Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// FindAllLocationsHandler
// @Summary Retrieve all Locations
// @Description Get a list of all Locations
// @Tags Locations
// @Accept json
// @Produce json
// @Success 200 {array} location_usecase.FindAllLocationsOutputDTO
// @Router /location [get]
func (h *LocationHandlers) FindAllLocationsHandler(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "limit", c.DefaultQuery("limit", "20"))
	ctx = context.WithValue(ctx, "offset", c.DefaultQuery("offset", "0"))
	res, err := location_usecase.NewFindAllLocationsUseCase(h.LocationRepository).Execute(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateLocationHandler
// @Summary Update a Location
// @Description Update a specific Location entity
// @Tags Locations
// @Accept json
// @Produce json
// @Param location_id path string true "Location ID"
// @Param input body location_usecase.UpdateLocationInputDTO true "Location entity to update"
// @Success 200 {object} location_usecase.UpdateLocationOutputDTO
// @Router /location/{location_id} [put]
func (h *LocationHandlers) UpdateLocationHandler(c *gin.Context) {
	var input location_usecase.UpdateLocationInputDTO
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	locationId, err := uuid.Parse(c.Param("location_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	input.LocationId = locationId
	ctx := context.Background()
	res, err := location_usecase.NewUpdateLocationUseCase(h.LocationRepository).Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteLocationHandler
// @Summary Delete a Location
// @Description Remove a specific Location from the system
// @Tags Locations
// @Accept json
// @Produce json
// @Param location_id path string true "Location ID"
// @Success 200 {string} string "Location deleted successfully"
// @Router /location/{location_id} [delete]
func (h *LocationHandlers) DeleteLocationHandler(c *gin.Context) {
	var input location_usecase.DeleteLocationInputDTO
	locationId, err := uuid.Parse(c.Param("location_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	input.LocationId = locationId
	ctx := context.Background()
	err = location_usecase.NewDeleteLocationUseCase(h.LocationRepository).Execute(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Location deleted successfully"})
}
