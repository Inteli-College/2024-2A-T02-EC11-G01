package entity

type Location struct {
	LocationId  *string `json:"location_id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        *string `json:"name"`
	CoordinateX *string `json:"coordinate_x"`
	CoordinateY *string `json:"coordinate_y"`
}

func NewLocation(name string, coordinateX string, coordinateY string) *Location {
	return &Location{
		Name:        &name,
		CoordinateX: &coordinateX,
		CoordinateY: &coordinateY,
	}
}
