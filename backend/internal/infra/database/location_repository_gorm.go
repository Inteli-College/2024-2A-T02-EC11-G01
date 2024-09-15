package database

import (
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

func (r *LocationRepositoryGorm) CreateLocation(input *entity.Location) (*entity.Location, error) {
	err := r.Db.Create(&input).Error
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *LocationRepositoryGorm) FindLocationById(id uuid.UUID) (*entity.Location, error) {
	var prediction entity.Location
	err := r.Db.First(&prediction, id).Error
	if err != nil {
		return nil, err
	}
	return &prediction, nil
}

func (r *LocationRepositoryGorm) FindAllLocations() ([]*entity.Location, error) {
	var locations []*entity.Location
	err := r.Db.Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (r *LocationRepositoryGorm) UpdateLocation(input *entity.Location) (*entity.Location, error) {
	var location entity.Location
	err := r.Db.First(&location, "id = ?", input.Id).Error
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

	res := r.Db.Save(location)
	if res.Error != nil {
		return nil, res.Error
	}
	return &location, nil
}

func (r *LocationRepositoryGorm) DeleteLocation(id uuid.UUID) error {
	err := r.Db.Delete(&entity.Location{}, "id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}
