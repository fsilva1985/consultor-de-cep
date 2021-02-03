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

	var slices [][]model.Address = createAddressChunk(addresses, 100)

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

// createAddressChunk returns [][]model.Address
func createAddressChunk(arr []model.Address, limit int) [][]model.Address {
	var slices [][]model.Address
	for i := 0; i < len(arr); i += limit {
		slices = append(slices, arr[i:i+limit])
	}

	return slices
}
