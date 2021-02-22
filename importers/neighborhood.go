package importers

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/djimenez/iconv-go"
	"github.com/fsilva1985/consultor-de-cep/entities"
	"github.com/fsilva1985/consultor-de-cep/parsers"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Neighborhood returns void
func Neighborhood(file io.ReadCloser, db *gorm.DB) {
	var neighborhoods []entities.Neighborhood

	i := 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		stringer, _ := iconv.ConvertString(scanner.Text(), "windows-1252", "utf-8")

		row := strings.Split(stringer, "@")

		neighborhoods = append(neighborhoods, entities.Neighborhood{
			ID:     parsers.StringToUint(row[0]),
			CityID: parsers.StringToUint(row[2]),
			Name:   row[3],
		})

		if i == 1000 {
			upsertNeighborhood(neighborhoods, db)
			i = 1
			neighborhoods = nil
		}

		i++
	}

	upsertNeighborhood(neighborhoods, db)

	fmt.Println("Bairros importados com sucesso")
}

func upsertNeighborhood(data []entities.Neighborhood, db *gorm.DB) {
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"city_id", "name"}),
	}).Create(data)
}
