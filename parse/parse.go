package parse

import (
	"strconv"

	"github.com/fsilva1985/consultor-de-cep/model"
)

// StringToUint returns uint
func StringToUint(value string) uint {
	number, _ := strconv.ParseUint(value, 10, 32)

	return uint(number)
}

// CreateCityChunk returns [][]model.City
func CreateCityChunk(arr []model.City, limit int) [][]model.City {
	var slices [][]model.City
	for i := 0; i < len(arr); i += limit {
		slices = append(slices, arr[i:i+limit])
	}

	return slices
}

// CreateNeighborhoodChunk returns [][]model.Neighborhood
func CreateNeighborhoodChunk(arr []model.Neighborhood, limit int) [][]model.Neighborhood {
	var slices [][]model.Neighborhood
	for i := 0; i < len(arr); i += limit {
		slices = append(slices, arr[i:i+limit])
	}

	return slices
}

// CreateAddressChunk returns [][]model.Address
func CreateAddressChunk(arr []model.Address, limit int) [][]model.Address {
	var slices [][]model.Address
	for i := 0; i < len(arr); i += limit {
		slices = append(slices, arr[i:i+limit])
	}

	return slices
}
