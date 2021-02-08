package importer

import (
	"bufio"
	"io"
	"strings"
	"sync"

	"github.com/djimenez/iconv-go"
	"github.com/fsilva1985/consultor-de-cep/model"
	"github.com/fsilva1985/consultor-de-cep/parse"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Address returns void
func Address(file io.ReadCloser, db *gorm.DB, stateCode string, wg *sync.WaitGroup) {

	defer wg.Done()

	var addresses []model.Address

	i := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringer, _ := iconv.ConvertString(scanner.Text(), "windows-1252", "utf-8")

		row := strings.Split(stringer, "@")

		addresses = append(addresses, model.Address{
			ID:             parse.StringToUint(row[0]),
			NeighborhoodID: parse.StringToUint(row[3]),
			Zipcode:        row[7],
			Type:           row[8],
			Name:           row[5],
		})

		if i == 1000 {
			upsertAddress(addresses, db)
			i = 1
			addresses = nil
		}

		i++
	}

	upsertAddress(addresses, db)
}

func upsertAddress(data []model.Address, db *gorm.DB) {
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"neighborhood_id", "name"}),
	}).Create(data)
}
