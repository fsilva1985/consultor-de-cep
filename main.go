package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type City struct {
	Id        uint `gorm:"primaryKey"`
	StateCode string
	Name      string
}

func main() {

	pwd, _ := os.Getwd()

	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&City{})

	getContentFile(
		pwd+"/eDNE_Basico.zip",
		"Delimitado/LOG_LOCALIDADE.TXT",
		db,
	)
}

// getContentFile returns []byte
func getContentFile(zipPath string, filepath string, db *gorm.DB) {

	read, _ := zip.OpenReader(zipPath)

	defer read.Close()

	for _, file := range read.File {
		if strings.Compare(file.Name, filepath) != 0 {
			continue
		}

		buffer, _ := file.Open()

		defer buffer.Close()

		scanner := bufio.NewScanner(buffer)

		for scanner.Scan() {
			row := strings.Split(scanner.Text(), "@")

			number, _ := strconv.ParseUint(row[0], 10, 32)
			id := uint(number)

			db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "id"}},
				DoUpdates: clause.AssignmentColumns([]string{"state_code", "name"}),
			}).Create(&City{
				Id:        id,
				StateCode: row[1],
				Name:      row[2],
			})
		}
	}

	fmt.Println("Cidades importadas.")
}
