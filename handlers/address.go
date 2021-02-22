package handlers

import (
	"net/http"

	"github.com/Nhanderu/brdoc"
	"github.com/fsilva1985/consultor-de-cep/entities"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllAddress(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	type Response struct {
		Type         string `json:"type"`
		Name         string `json:"name"`
		Neighborhood string `json:"neighborhood"`
		City         string `json:"city"`
		State        string `json:"state"`
	}

	vars := mux.Vars(r)

	if !brdoc.IsCEP(vars["zipcode"]) {
		ResponseJson(w, 422, map[string]string{"message": "zipcode is not valid."})

		return
	}

	var address entities.Address

	err := db.Preload("Neighborhood.City.State").First(&address, "Zipcode = ?", vars["zipcode"]).Error

	if err != nil {
		ResponseJson(w, 422, map[string]string{"message": "zipcode is not found."})
	}

	m := Response{
		Type:         address.Type,
		Name:         address.Name,
		Neighborhood: address.Neighborhood.Name,
		City:         address.Neighborhood.City.Name,
		State:        address.Neighborhood.City.State.Name,
	}

	ResponseJson(w, http.StatusOK, map[string]interface{}{"data": m})
}
