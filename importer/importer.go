package importer

import (
	"bufio"
	"fmt"
	"runtime"
	"strings"

	"github.com/fsilva1985/consultor-de-cep/console"
	"github.com/fsilva1985/consultor-de-cep/model"
	"github.com/fsilva1985/consultor-de-cep/parse"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// City returns void
func City(buffer *bufio.Scanner, db *gorm.DB) {
	var cities []model.City

	for buffer.Scan() {
		row := strings.Split(buffer.Text(), "@")

		var state model.State

		db.First(&state, "code = ?", row[1])

		cities = append(cities, model.City{
			Id:      parse.StringToUint(row[0]),
			StateId: state.Id,
			Name:    row[2],
		})
	}

	var slices [][]model.City = parse.CreateCityChunk(cities, 100)

	done := make(chan string)
	go upsertCity(slices, db, done)
	fmt.Println(<-done)
}

func upsertCity(slices [][]model.City, db *gorm.DB, done chan string) {
	for _, slice := range slices {
		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"state_id", "name"}),
		}).Create(slice)
	}
	done <- console.GetDoneMessage("Cidades importados")
}

// Neighborhood returns void
func Neighborhood(buffer *bufio.Scanner, db *gorm.DB) {
	var neighborhoods []model.Neighborhood
	for buffer.Scan() {
		row := strings.Split(buffer.Text(), "@")

		neighborhoods = append(neighborhoods, model.Neighborhood{
			Id:     parse.StringToUint(row[0]),
			CityId: parse.StringToUint(row[2]),
			Name:   row[3],
		})

	}

	var slices [][]model.Neighborhood = parse.CreateNeighborhoodChunk(neighborhoods, 100)

	done := make(chan string)
	go upsertNeighborhood(slices, db, done)
	fmt.Println(<-done)
}

func upsertNeighborhood(slices [][]model.Neighborhood, db *gorm.DB, done chan string) {
	for _, slice := range slices {
		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"city_id", "name"}),
		}).Create(slice)
	}
	done <- console.GetDoneMessage("Bairros importados")
}

// Address returns void
func Address(buffer *bufio.Scanner, db *gorm.DB, stateCode string) {
	var addresses []model.Address
	for buffer.Scan() {
		row := strings.Split(buffer.Text(), "@")

		addresses = append(addresses, model.Address{
			Id:             parse.StringToUint(row[0]),
			NeighborhoodId: parse.StringToUint(row[3]),
			Zipcode:        row[7],
			Type:           row[8],
			Name:           row[5],
		})
	}

	var slices [][]model.Address = parse.CreateAddressChunk(addresses, 100)

	done := make(chan string)
	go upsertAddress(slices, db, stateCode, done)
	fmt.Println(<-done)
}

func upsertAddress(slices [][]model.Address, db *gorm.DB, stateCode string, done chan string) {
	for _, slice := range slices {
		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"neighborhood_id", "name"}),
		}).Create(slice)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	done <- console.GetDoneMessage("Endereços do Estado " + stateCode + " importados")
}
