package dto

type CreateLocationInputDTO struct {
	Name        *string `json:"name"`
	CoordinateX *string `json:"coordinate_x"`
	CoordinateY *string `json:"coordinate_y"`
}

type LocationOutputDTO struct {
	LocationId  *string `json:"location_id"`
	Name        *string `json:"name"`
	CoordinateX *string `json:"coordinate_x"`
	CoordinateY *string `json:"coordinate_y"`
}
