package entity

type Location struct {
	LocationId  *string "json:\"location_id\""
	Name        *string "json:\"name\""
	CoordinateX *string "json:\"coordinate_x\""
	CoordinateY *string "json:\"coordinate_y\""
}

func NewLocation(locationId string, name string, coordinateX string, coordinateY string) *Location {
	return &Location{
		LocationId:  &locationId,
		Name:        &name,
		CoordinateX: &coordinateX,
		CoordinateY: &coordinateY,
	}
}

func (p *Location) Validate() error {
	return nil
}
