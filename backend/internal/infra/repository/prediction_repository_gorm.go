package repository

import (
	"context"
	"errors"
	"log"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/mappers"
	"gorm.io/gorm"
)

type PredictionRepository struct {
	db *gorm.DB
}

func NewPredictionRepository(db *gorm.DB) *PredictionRepository {
	log.Printf("Prediction DB address: %v\n", db)
	return &PredictionRepository{
		db: db,
	}
}

func (r *PredictionRepository) Create(ctx context.Context, predictionDTO *dto.CreatePredictionInputDTO) (*dto.PredictionDTO, error) {
	prediction := entity.NewPrediction(predictionDTO.RawImagePath, predictionDTO.OutputImagePath, predictionDTO.Output, predictionDTO.LocationId)

	if err := r.db.WithContext(ctx).Create(prediction).Error; err != nil {
		return nil, err
	}

	return mappers.MapPredictionEntityToDTO(prediction), nil
}

func (r *PredictionRepository) GetByID(ctx context.Context, id *string) (*dto.PredictionDTO, error) {
	var prediction entity.Prediction

	if err := r.db.WithContext(ctx).First(&prediction, "prediction_id = ?", *id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // TODO: return personalized errors
		}
		return nil, err
	}

	return mappers.MapPredictionEntityToDTO(&prediction), nil
}

func (r *PredictionRepository) GetAllByLocationID(ctx context.Context, locationID *string, limit *int, offset *int, orders ...string) ([]*dto.PredictionDTO, error) {
	var predictions []*entity.Prediction
	query := r.db.WithContext(ctx)

	query = r.addOrderToQuery(query, orders)
	query = r.addLimitToQuery(query, limit)
	query = r.addOffsetToQuery(query, offset)

	if err := query.Where("location_id = ?", *locationID).Find(&predictions).Error; err != nil {
		return nil, err // TODO: return personalized errors
	}

	predictionsOutputs := make([]*dto.PredictionDTO, 0, len(predictions))
	for _, prediction := range predictions {
		predictionsOutputs = append(predictionsOutputs, mappers.MapPredictionEntityToDTO(prediction))
	}

	return predictionsOutputs, nil
}

func (r *PredictionRepository) GetAll(ctx context.Context, limit *int, offset *int, orders ...string) ([]*dto.PredictionDTO, error) {
	var predictions []*entity.Prediction
	query := r.db.WithContext(ctx)

	query = r.addOrderToQuery(query, orders)
	query = r.addLimitToQuery(query, limit)
	query = r.addOffsetToQuery(query, offset)

	if err := query.Find(&predictions).Error; err != nil {
		return nil, err
	}

	predictionsOutputs := make([]*dto.PredictionDTO, 0, len(predictions))
	for _, prediction := range predictions {
		predictionsOutputs = append(predictionsOutputs, mappers.MapPredictionEntityToDTO(prediction))
	}

	return predictionsOutputs, nil
}

func (r *PredictionRepository) addOrderToQuery(query *gorm.DB, orders []string) *gorm.DB {
	if orders != nil && len(orders) > 0 {
		for _, order := range orders {
			query = query.Order(order)
		}
		return query
	}
	return query.Order("created_at DESC")
}

func (r *PredictionRepository) addLimitToQuery(query *gorm.DB, limit *int) *gorm.DB {
	if limit != nil {
		return query.Limit(*limit)
	}

	return query
}

func (r *PredictionRepository) addOffsetToQuery(query *gorm.DB, offset *int) *gorm.DB {
	if offset != nil {
		return query.Offset(*offset)
	}

	return query
}
