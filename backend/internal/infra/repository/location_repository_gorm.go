package repository

import (
	"context"
	"strconv"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LocationRepositoryGorm struct {
	Db *gorm.DB
}


func NewLocationRepositoryGorm(db *gorm.DB) *LocationRepositoryGorm {
	return &LocationRepositoryGorm{
		Db: db,
	}
}

func (r *LocationRepositoryGorm) CreateLocation(ctx context.Context, input *entity.Location) (*entity.Location, error) {
	err := r.Db.WithContext(ctx).Create(&input).Error
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *LocationRepositoryGorm) FindLocationById(ctx context.Context, id uuid.UUID) (*entity.Location, error) {
	var prediction entity.Location
	err := r.Db.WithContext(ctx).Preload("Predictions").First(&prediction, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrLocationNotFound
		}
		return nil, err
	}
	return &prediction, nil
}

func (r *LocationRepositoryGorm) FindAllLocations(ctx context.Context) ([]*entity.Location, error) {
	var locations []*entity.Location
	r.pagination(ctx, r.Db)
	err := r.Db.WithContext(ctx).Preload("Predictions").Find(&locations).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrLocationNotFound
		}
		return nil, err
	}
	return locations, nil
}

func (r *LocationRepositoryGorm) UpdateLocation(ctx context.Context, input *entity.Location) (*entity.Location, error) {
	var location entity.Location
	err := r.Db.WithContext(ctx).First(&location, "location_id = ?", input.LocationId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrLocationNotFound
		}
		return nil, err
	}

	location.Name = input.Name
	location.Latitude = input.Latitude
	location.Longitude = input.Longitude
	location.UpdatedAt = input.UpdatedAt

	res := r.Db.WithContext(ctx).Save(location)
	if res.Error != nil {
		return nil, res.Error
	}
	return &location, nil
}

func (r *LocationRepositoryGorm) DeleteLocation(ctx context.Context, locationId uuid.UUID) error {
	err := r.Db.WithContext(ctx).Delete(&entity.Location{}, "location_id = ?", locationId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.ErrLocationNotFound
		}
		return err
	}
	return nil
}

func (r *LocationRepositoryGorm) pagination(ctx context.Context, tx *gorm.DB) (*gorm.DB, error) {
	limit, err := strconv.Atoi(ctx.Value("limit").(string))
	if err != nil {
		return nil, err
	}
	offset, err := strconv.Atoi(ctx.Value("offset").(string))
	if err != nil {
		return nil, err
	}
	tx.Limit(limit)
	tx.Offset(offset)
	return tx, nil
}
