package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/fsilva1985/consultor-de-cep/handler"
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
		vars := mux.Vars(r)
		var address model.Address

		a.DB.Preload("Neighborhood.City.State").First(&address, "Zipcode = ?", vars["zipcode"])

		handler.RespondWithData(w, http.StatusOK, address)
	}).Methods("GET")
}

// Run app
func (a *App) Run(addr string) {
	fmt.Println("Server Started")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
