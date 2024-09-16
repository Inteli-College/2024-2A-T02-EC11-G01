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

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	log.Printf("Location DB address: %p\n", db)
	return &LocationRepository{db: db}
}

func (r *LocationRepository) CreateLocation(ctx context.Context, input *dto.CreateLocationInputDTO) (*dto.LocationOutputDTO, error) {
	location := entity.NewLocation(*input.Name, *input.CoordinateX, *input.CoordinateY)
	result := r.db.WithContext(ctx).Create(location)
	if result.Error != nil {
		return nil, result.Error
	}

	return mappers.MapLocationEntityToDTO(location), nil
}

func (r *LocationRepository) GetLocationById(ctx context.Context, locationId string) (*dto.LocationOutputDTO, error) {
	var location entity.Location
	result := r.db.WithContext(ctx).First(&location, "location_id = ?", locationId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("location not found")
	}

	return mappers.MapLocationEntityToDTO(&location), nil
}

func (r *LocationRepository) UpdateLocation(ctx context.Context, locationId string, input *dto.CreateLocationInputDTO) (*dto.LocationOutputDTO, error) {
	var location entity.Location
	result := r.db.WithContext(ctx).First(&location, "location_id = ?", locationId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("location not found")
	}

	if input.Name != nil {
		location.Name = input.Name
	}
	if input.CoordinateX != nil {
		location.CoordinateX = input.CoordinateX
	}
	if input.CoordinateY != nil {
		location.CoordinateY = input.CoordinateY
	}

	if err := r.db.WithContext(ctx).Model(&location).Where("location_id = ?", locationId).Updates(location).Error; err != nil {
		return nil, err
	}

	return mappers.MapLocationEntityToDTO(&location), nil
}

func (r *LocationRepository) DeleteLocation(ctx context.Context, locationId string) error {
	result := r.db.WithContext(ctx).Delete(&entity.Location{}, "location_id = ?", locationId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *LocationRepository) GetAllLocations(ctx context.Context) ([]*dto.LocationOutputDTO, error) {
	var locations []*entity.Location
	result := r.db.WithContext(ctx).Find(&locations)
	if result.Error != nil {
		return nil, result.Error
	}

	locationOutputs := make([]*dto.LocationOutputDTO, 0, len(locations))
	for _, location := range locations {
		locationOutputs = append(locationOutputs, mappers.MapLocationEntityToDTO(location))
	}

	return locationOutputs, nil
}
