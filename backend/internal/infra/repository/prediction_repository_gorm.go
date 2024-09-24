package repository

import (
	"context"
	"errors"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PredictionRepositoryGorm struct {
	db *gorm.DB
}

func NewPredictionRepositoryGorm(db *gorm.DB) *PredictionRepositoryGorm {
	return &PredictionRepositoryGorm{
		db: db,
	}
}

func (r *PredictionRepositoryGorm) CreatePrediction(ctx context.Context, input *entity.Prediction) (*entity.Prediction, error) {
	if err := r.db.WithContext(ctx).Create(input).Error; err != nil {
		return nil, err
	}

	return input, nil
}

func (r *PredictionRepositoryGorm) FindPredictionById(ctx context.Context, id *uuid.UUID) (*entity.Prediction, error) {
	var prediction entity.Prediction

	if err := r.db.WithContext(ctx).First(&prediction, "prediction_id = ?", *id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err // TODO: return personalized errors
		}
		return nil, err
	}

	return &prediction, nil
}

func (r *PredictionRepositoryGorm) FindAllPredictionsByLocationID(ctx context.Context, locationID *uuid.UUID, limit *int, offset *int, orders ...string) ([]*entity.Prediction, error) {
	var predictions []*entity.Prediction
	query := r.db.WithContext(ctx)

	query = r.addOrderToQuery(query, orders)
	query = r.addLimitToQuery(query, limit)
	query = r.addOffsetToQuery(query, offset)

	if err := query.Where("location_id = ?", *locationID).Find(&predictions).Error; err != nil {
		return nil, err // TODO: return personalized errors
	}

	return predictions, nil
}

func (r *PredictionRepositoryGorm) FindAllPredictions(ctx context.Context, limit *int, offset *int, orders ...string) ([]*entity.Prediction, error) {
	var predictions []*entity.Prediction
	query := r.db.WithContext(ctx)

	query = r.addOrderToQuery(query, orders)
	query = r.addLimitToQuery(query, limit)
	query = r.addOffsetToQuery(query, offset)

	if err := query.Find(&predictions).Error; err != nil {
		return nil, err
	}

	return predictions, nil
}

func (r *PredictionRepositoryGorm) addOrderToQuery(query *gorm.DB, orders []string) *gorm.DB {
	if orders != nil && len(orders) > 0 {
		for _, order := range orders {
			query = query.Order(order)
		}
		return query
	}
	return query.Order("created_at DESC")
}

func (r *PredictionRepositoryGorm) addLimitToQuery(query *gorm.DB, limit *int) *gorm.DB {
	if limit != nil {
		return query.Limit(*limit)
	}

	return query
}

func (r *PredictionRepositoryGorm) addOffsetToQuery(query *gorm.DB, offset *int) *gorm.DB {
	if offset != nil {
		return query.Offset(*offset)
	}

	return query
}
