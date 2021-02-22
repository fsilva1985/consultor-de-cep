package entities

// Neighborhood returns collection
type Neighborhood struct {
	ID     uint
	CityID uint
	Name   string
	City   City
}
