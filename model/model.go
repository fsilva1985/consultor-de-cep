package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// State returns collection
type State struct {
	Id   uint
	Code string
	Name string
}

// City returns collection
type City struct {
	Id      uint
	StateId uint
	Name    string
}

// Neighborhood returns collection
type Neighborhood struct {
	Id     uint
	CityId uint
	Name   string
}

// Address returns collection
type Address struct {
	Id             uint
	NeighborhoodId uint
	Zipcode        string
	Type           string
	Name           string
}

// GetStates returns collection
func GetStates() []State {
	var states = []State{
		{Id: 1, Code: "AC", Name: "Acre"},
		{Id: 2, Code: "AL", Name: "Alagoas"},
		{Id: 3, Code: "AP", Name: "Amapá"},
		{Id: 4, Code: "AM", Name: "Amazonas"},
		{Id: 5, Code: "BA", Name: "Bahia"},
		{Id: 6, Code: "CE", Name: "Ceará"},
		{Id: 7, Code: "DF", Name: "Distrito Federal"},
		{Id: 8, Code: "ES", Name: "Espírito Santo"},
		{Id: 9, Code: "GO", Name: "Goiás"},
		{Id: 10, Code: "MA", Name: "Maranhão"},
		{Id: 11, Code: "MT", Name: "Mato Grosso"},
		{Id: 12, Code: "MS", Name: "Mato Grosso do Sul"},
		{Id: 13, Code: "MG", Name: "Minas Gerais"},
		{Id: 14, Code: "PA", Name: "Pará"},
		{Id: 15, Code: "PB", Name: "Paraíba"},
		{Id: 16, Code: "PR", Name: "Paraná"},
		{Id: 17, Code: "PE", Name: "Pernambuco"},
		{Id: 18, Code: "PI", Name: "Piauí"},
		{Id: 19, Code: "RJ", Name: "Rio de Janeiro"},
		{Id: 20, Code: "RN", Name: "Rio Grande do Norte"},
		{Id: 21, Code: "RS", Name: "Rio Grande do Sul"},
		{Id: 22, Code: "RO", Name: "Rondônia"},
		{Id: 23, Code: "RR", Name: "Roraima"},
		{Id: 24, Code: "SC", Name: "Santa Catarina"},
		{Id: 25, Code: "SP", Name: "São Paulo"},
		{Id: 26, Code: "SE", Name: "Sergipe"},
		{Id: 27, Code: "TO", Name: "Tocantins"},
	}

	return states
}

// Init returns *gorm.DB
func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("model/database.sqlite"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&State{}, &City{}, &Neighborhood{}, &Address{})

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"code", "name"}),
	}).Create(GetStates())

	return db
}
