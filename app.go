package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nhanderu/brdoc"
	"github.com/fsilva1985/consultor-de-cep/httpResponse"
	"github.com/fsilva1985/consultor-de-cep/importer"
	"github.com/fsilva1985/consultor-de-cep/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// App instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize instances
func (a *App) Initialize() {
	a.DB = model.Initialize()

	importer.Initialize(a.DB)

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/api/address/{zipcode}", func(w http.ResponseWriter, r *http.Request) {
		type Response struct {
			Type         string `json:"type"`
			Name         string `json:"name"`
			Neighborhood string `json:"neighborhood"`
			City         string `json:"city"`
			State        string `json:"state"`
		}

		vars := mux.Vars(r)

		if !brdoc.IsCEP(vars["zipcode"]) {
			httpResponse.ResponseMessage(w, 422, "zipcode is not valid.")

			return
		}

		var address model.Address

		err := a.DB.Preload("Neighborhood.City.State").First(&address, "Zipcode = ?", vars["zipcode"]).Error

		if err != nil {
			httpResponse.ResponseMessage(w, 422, "zipcode is not found.")

			return
		}

		m := Response{
			Type:         address.Type,
			Name:         address.Name,
			Neighborhood: address.Neighborhood.Name,
			City:         address.Neighborhood.City.Name,
			State:        address.Neighborhood.City.State.Name,
		}

		httpResponse.ResponseData(w, http.StatusOK, m)
	}).Methods("GET")
}

// Run app
func (a *App) Run(addr string) {
	fmt.Println("Server Started")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
