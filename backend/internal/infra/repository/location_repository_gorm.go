package repository

import (
	"context"
	"errors"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LocationRepositoryGorm struct {
	db *gorm.DB
}

func NewLocationRepositoryGorm(db *gorm.DB) *LocationRepositoryGorm {
	return &LocationRepositoryGorm{db: db}
}

func (r *LocationRepositoryGorm) CreateLocation(ctx context.Context, input *entity.Location) (*entity.Location, error) {
	result := r.db.WithContext(ctx).Create(input)
	if result.Error != nil {
		return nil, result.Error
	}

	return input, nil
}

func (r *LocationRepositoryGorm) GetLocationById(ctx context.Context, locationId *uuid.UUID) (*entity.Location, error) {
	var location entity.Location
	result := r.db.WithContext(ctx).First(&location, "location_id = ?", *locationId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("location not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &location, nil
}

func (r *LocationRepositoryGorm) UpdateLocation(ctx context.Context, input *entity.Location) (*entity.Location, error) {
	res := r.db.WithContext(ctx).Model(&input).Where("location_id = ?", input.LocationId).Updates(input)

	if res.Error != nil {
		return nil, res.Error
	}
	return input, nil
}

func (r *LocationRepositoryGorm) DeleteLocation(ctx context.Context, locationId *uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&entity.Location{}, "location_id = ?", *locationId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *LocationRepositoryGorm) GetAllLocations(ctx context.Context) ([]*entity.Location, error) {
	var locations []*entity.Location
	result := r.db.WithContext(ctx).Find(&locations)
	if result.Error != nil {
		return nil, result.Error
	}

	return locations, nil
}
