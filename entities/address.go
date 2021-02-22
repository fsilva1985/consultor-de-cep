package entities

// Address returns collection
type Address struct {
	ID             uint
	NeighborhoodID uint
	Zipcode        string
	Type           string
	Name           string
	Neighborhood   Neighborhood
}
