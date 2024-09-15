package database

import (
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

func (r *PredictionRepositoryGorm) CreatePrediction(input *entity.Prediction) (*entity.Prediction, error) {
	err := r.Db.Create(&input).Error
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *PredictionRepositoryGorm) FindPredictionById(id uuid.UUID) (*entity.Prediction, error) {
	var prediction entity.Prediction
	err := r.Db.Preload("Locations").First(&prediction, id).Error
	if err != nil {
		return nil, err
	}
	return &prediction, nil
}

func (r *PredictionRepositoryGorm) FindAllPredictions() ([]*entity.Prediction, error) {
	var predictions []*entity.Prediction
	err := r.Db.Preload("Locations").Find(&predictions).Error
	if err != nil {
		return nil, err
	}
	return predictions, nil
}

func (r *PredictionRepositoryGorm) UpdatePrediction(input *entity.Prediction) (*entity.Prediction, error) {
	var prediction entity.Prediction
	err := r.Db.First(&prediction, "id = ?", input.Id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrPredictionNotFound
		}
		return nil, err
	}

	prediction.RawImage = input.RawImage
	prediction.AnnotatedImage = input.AnnotatedImage
	prediction.Detections = input.Detections
	prediction.LocationId = input.LocationId
	prediction.UpdatedAt = input.UpdatedAt

	res := r.Db.Save(prediction)
	if res.Error != nil {
		return nil, res.Error
	}
	return &prediction, nil
}

func (r *PredictionRepositoryGorm) DeletePrediction(id uuid.UUID) error {
	err := r.Db.Delete(&entity.Prediction{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
