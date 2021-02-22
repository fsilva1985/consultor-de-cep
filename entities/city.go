package entities

// City returns collection
type City struct {
	ID      uint
	StateID uint
	Name    string
	State   State
}
