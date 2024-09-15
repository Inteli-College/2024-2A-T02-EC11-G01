package handler

import (
	"net/http"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/usecase/prediction_usecase"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PredictionHandlers struct {
	EventDispatcher        events.EventDispatcherInterface
	PredictionRepository   entity.PredictionRepository
	PredictionCreatedEvent events.EventInterface
}

func NewPredictionHandlers(
	eventDispatcher events.EventDispatcherInterface,
	predictionRepository entity.PredictionRepository,
	predictionCreatedEvent events.EventInterface,
) *PredictionHandlers {
	return &PredictionHandlers{
		EventDispatcher:        eventDispatcher,
		PredictionRepository:   predictionRepository,
		PredictionCreatedEvent: predictionCreatedEvent,
	}
}

func (h *PredictionHandlers) CreatePredictionHandler(c *gin.Context) {
	var input prediction_usecase.CreatePredictionInputDTO
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := prediction_usecase.NewCreatePredictionUseCase(
		h.PredictionCreatedEvent,
		h.PredictionRepository,
		h.EventDispatcher,
	).Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *PredictionHandlers) FindPredictionByIdHandler(c *gin.Context) {
	var input prediction_usecase.FindPredictionByIdInputDTO
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := prediction_usecase.NewFindPredictionByIdUseCase(h.PredictionRepository).Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *PredictionHandlers) FindAllPredictionsHandler(c *gin.Context) {
	res, err := prediction_usecase.NewFindAllPredictionsUseCase(h.PredictionRepository).Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *PredictionHandlers) UpdatePredictionHandler(c *gin.Context) {
	var input prediction_usecase.UpdatePredictionInputDTO
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	input.Id = id
	res, err := prediction_usecase.NewUpdatePredictionUseCase(h.PredictionRepository).Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *PredictionHandlers) DeletePredictionHandler(c *gin.Context) {
	var input prediction_usecase.DeletePredictionInputDTO
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	input.Id = id
	err = prediction_usecase.NewDeletePredictionUseCase(h.PredictionRepository).Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Prediction deleted successfully"})
}
