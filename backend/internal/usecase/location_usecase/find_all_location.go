package location_usecase

import "github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"

type FindAllLocationsOutputDTO []*FindLocationOutputDTO

type FindAllLocationsUseCase struct {
	LocationRepository entity.LocationRepository
}

func NewFindAllLocationsUseCase(locationRepository entity.LocationRepository) *FindAllLocationsUseCase {
	return &FindAllLocationsUseCase{
		LocationRepository: locationRepository,
	}
}

func (u *FindAllLocationsUseCase) Execute() (*FindAllLocationsOutputDTO, error) {
	res, err := u.LocationRepository.FindAllLocations()
	if err != nil {
		return nil, err
	}
	output := make(FindAllLocationsOutputDTO, len(res))
	for i, location := range res {
		output[i] = &FindLocationOutputDTO{
			Id:        location.Id,
			Name:      location.Name,
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
			CreatedAt: location.CreatedAt,
			UpdatedAt: location.UpdatedAt,
		}
	}
	return &output, nil
}
