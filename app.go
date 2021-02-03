package main

import (
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
	a.Router.HandleFunc("/address/{zipcode}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var address model.Address

		a.DB.First(&address, "Zipcode = ?", vars["zipcode"])

		handler.RespondWithJSON(w, http.StatusOK, address)
	}).Methods("GET")
}

// Run app
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
