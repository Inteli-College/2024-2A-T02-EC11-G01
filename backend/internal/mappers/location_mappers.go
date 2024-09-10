package mappers

import (
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
)

func MapLocationEntityToDTO(location *entity.Location) *dto.LocationOutputDTO {
	return &dto.LocationOutputDTO{
		LocationId:  location.LocationId,
		Name:        location.Name,
		CoordinateX: location.CoordinateX,
		CoordinateY: location.CoordinateY,
	}
}
