package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/infra/repository"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/prediction_usecase"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/gin-gonic/gin"
)

type PredictionHandler struct {
	createPredictionUsecase               *prediction_usecase.CreatePredictionUseCase
	findPredictionByIdUsecase             *prediction_usecase.FindPredictionByIdUseCase
	findAllPredictionsUsecase             *prediction_usecase.FindAllPredictionsUseCase
	findAllPredictionsByLocationIdUsecase *prediction_usecase.FindAllPredictionsByLocationIdUsecase
}

func NewPredictionHandler(locationRepo *repository.PredictionRepositoryGorm, eventDispatcher events.EventDispatcherInterface, predictionCreatedEvent events.EventInterface) *PredictionHandler {
	handler := &PredictionHandler{
		createPredictionUsecase:               prediction_usecase.NewCreatePredictionUseCase(predictionCreatedEvent, locationRepo, eventDispatcher),
		findPredictionByIdUsecase:             prediction_usecase.NewFindPredictionByIdUseCase(locationRepo),
		findAllPredictionsUsecase:             prediction_usecase.NewFindAllPredictionsUseCase(locationRepo),
		findAllPredictionsByLocationIdUsecase: prediction_usecase.NewFindAllPredictionsByLocationIdUsecase(locationRepo),
	}

	return handler
}

// CreatePrediction godoc
// @Summary Create a new prediction
// @Description Create a new prediction
// @Tags Predictions
// @Accept json
// @Produce json
// @Param prediction body dto.CreatePredictionInputDTO true "Create Prediction Input"
// @Success 201 {object} dto.PredictionDTO
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /predictions [post]
func (h *PredictionHandler) CreatePrediction(c *gin.Context) {
	var input dto.CreatePredictionInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	prediction, err := h.createPredictionUsecase.Execute(ctx, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, prediction)
}

// FindAllPredictions godoc
// @Summary Get all predictions
// @Description Get a list of all predictions
// @Tags Predictions
// @Accept json
// @Produce json
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset for pagination"
// @Success 200 {array} dto.PredictionDTO
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /predictions [get]
func (h *PredictionHandler) FindAllPredictions(c *gin.Context) {
	ctx := context.Background()

	limitStr := c.DefaultQuery("limit", "")
	offsetStr := c.DefaultQuery("offset", "")

	var limit *int
	var offset *int

	if limitStr != "" {
		limitVerified, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
		limit = &limitVerified
	}

	if offsetStr != "" {
		offsetVerified, err := strconv.Atoi(offsetStr) // valor padrão 0
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
		offset = &offsetVerified
	}

	predictions, err := h.findAllPredictionsUsecase.Execute(ctx, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, predictions)
}

// FindAllPredictionsByLocationId godoc
// @Summary Get predictions by location Id
// @Description Get predictions by location Id
// @Tags Predictions
// @Accept json
// @Produce json
// @Param location_id path string true "Location ID"
// @Param limit query int false "Limit the number of results"
// @Param offset query int false "Offset for pagination"
// @Success 200 {array} dto.PredictionDTO
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /predictions/location/{location_id} [get]
func (h *PredictionHandler) FindAllPredictionsByLocationId(c *gin.Context) {
	locationId := c.Param("location_id")
	ctx := context.Background()

	limitStr := c.DefaultQuery("limit", "")
	offsetStr := c.DefaultQuery("offset", "")

	var limit *int
	var offset *int

	if limitStr != "" {
		limitVerified, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
		limit = &limitVerified
	}

	if offsetStr != "" {
		offsetVerified, err := strconv.Atoi(offsetStr) // valor padrão 0
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
		offset = &offsetVerified
	}

	var input dto.FindPredictionByLocationIdInputDTO
	input.LocationId = locationId

	location, err := h.findAllPredictionsByLocationIdUsecase.Execute(ctx, &input, limit, offset)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, location)
}

// FindPredictionByPredictionId godoc
// @Summary Find prediction by id
// @Description Find prediction by id
// @Tags Predictions
// @Accept json
// @Produce json
// @Param id path string true "Prediction ID"
// @Success 200 {object} dto.PredictionDTO
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /predictions/{id} [get]
func (h *PredictionHandler) FindPredictionByPredictionId(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")

	var input dto.FindPredictionByIdInputDTO
	input.PredictionId = id

	prediction, err := h.findPredictionByIdUsecase.Execute(ctx, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prediction)
}
