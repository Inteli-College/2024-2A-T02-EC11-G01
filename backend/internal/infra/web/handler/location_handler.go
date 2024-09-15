package handler

import (
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
// @Router /locations [post]
func (h *LocationHandlers) CreateLocationHandler(c *gin.Context) {
	var input location_usecase.CreateLocationInputDTO
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := location_usecase.NewCreateLocationUseCase(
		h.LocationCreatedEvent,
		h.LocationRepository,
		h.EventDispatcher,
	).Execute(input)
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
// @Param id path string true "Location ID"
// @Success 200 {object} location_usecase.FindLocationOutputDTO
// @Router /locations/{id} [get]
func (h *LocationHandlers) FindLocationByIdHandler(c *gin.Context) {
	var input location_usecase.FindLocationByIdInputDTO
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := location_usecase.NewFindLocationByIdUseCase(h.LocationRepository).Execute(input)
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
// @Router /locations [get]
func (h *LocationHandlers) FindAllLocationsHandler(c *gin.Context) {
	res, err := location_usecase.NewFindAllLocationsUseCase(h.LocationRepository).Execute()
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
// @Param id path string true "Location ID"
// @Param input body location_usecase.UpdateLocationInputDTO true "Location entity to update"
// @Success 200 {object} location_usecase.UpdateLocationOutputDTO
// @Router /locations/{id} [put]
func (h *LocationHandlers) UpdateLocationHandler(c *gin.Context) {
	var input location_usecase.UpdateLocationInputDTO
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	input.Id = id
	res, err := location_usecase.NewUpdateLocationUseCase(h.LocationRepository).Execute(input)
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
// @Param id path string true "Location ID"
// @Success 200 {string} string "Location deleted successfully"
// @Router /locations/{id} [delete]
func (h *LocationHandlers) DeleteLocationHandler(c *gin.Context) {
	var input location_usecase.DeleteLocationInputDTO
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	input.Id = id
	err = location_usecase.NewDeleteLocationUseCase(h.LocationRepository).Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Location deleted successfully"})
}
