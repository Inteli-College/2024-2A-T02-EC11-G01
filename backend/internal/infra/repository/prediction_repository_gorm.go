package repository

import (
	"context"

	"strconv"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PredictionRepositoryGorm struct {
	Db *gorm.DB
}

func NewPredictionRepositoryGorm(db *gorm.DB) *PredictionRepositoryGorm {
	return &PredictionRepositoryGorm{
		Db: db,
	}
}

func (r *PredictionRepositoryGorm) CreatePrediction(ctx context.Context, input *entity.Prediction) (*entity.Prediction, error) {
	err := r.Db.WithContext(ctx).Create(&input).Error
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *PredictionRepositoryGorm) FindPredictionById(ctx context.Context, id uuid.UUID) (*entity.Prediction, error) {
	var prediction entity.Prediction
	err := r.Db.WithContext(ctx).First(&prediction, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrPredictionNotFound
		}
		return nil, err
	}
	return &prediction, nil
}

func (r *PredictionRepositoryGorm) FindAllPredictionsByLocationId(ctx context.Context, id uuid.UUID) ([]*entity.Prediction, error) {
	var predictions []*entity.Prediction
	r.pagination(ctx, r.Db)
	err := r.Db.WithContext(ctx).Where("location_id = ?", id).Find(&predictions).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrPredictionNotFound
		}
		return nil, err
	}
	return predictions, nil
}

func (r *PredictionRepositoryGorm) FindAllPredictions(ctx context.Context) ([]*entity.Prediction, error) {
	var predictions []*entity.Prediction
	r.pagination(ctx, r.Db)
	err := r.Db.WithContext(ctx).Find(&predictions).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrPredictionNotFound
		}
		return nil, err
	}
	return predictions, nil
}

func (r *PredictionRepositoryGorm) UpdatePrediction(ctx context.Context, input *entity.Prediction) (*entity.Prediction, error) {
	var prediction entity.Prediction
	err := r.Db.WithContext(ctx).First(&prediction, "prediction_id = ?", input.PredictionId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrPredictionNotFound
		}
		return nil, err
	}

	prediction.RawImagePath = input.RawImagePath
	prediction.AnnotatedImagePath = input.AnnotatedImagePath
	prediction.Detections = input.Detections
	prediction.LocationId = input.LocationId
	prediction.UpdatedAt = input.UpdatedAt

	res := r.Db.WithContext(ctx).Save(prediction)
	if res.Error != nil {
		return nil, res.Error
	}
	return &prediction, nil
}

func (r *PredictionRepositoryGorm) DeletePrediction(ctx context.Context, predictionId uuid.UUID) error {
	err := r.Db.WithContext(ctx).Delete(&entity.Prediction{}, "prediction_id = ?", predictionId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.ErrPredictionNotFound
		}
		return err
	}
	return nil
}

func (r *PredictionRepositoryGorm) pagination(ctx context.Context, tx *gorm.DB) (*gorm.DB, error) {
	limit, err := strconv.Atoi(ctx.Value("limit").(string))
	if err != nil {
		return nil, err
	}
	offset, err := strconv.Atoi(ctx.Value("offset").(string))
	if err != nil {
		return nil, err
	}
	tx.Order("created_at DESC")
	tx.Limit(limit)
	tx.Offset(offset)
	return tx, nil
}
