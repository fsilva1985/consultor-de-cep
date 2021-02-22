package main

import (
	"log"
	"net/http"

	"github.com/fsilva1985/consultor-de-cep/handlers"
	"github.com/fsilva1985/consultor-de-cep/importers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (app *App) Initialize() {
	app.DB = getDBInstance()

	importers.Initialize(app.DB)

	app.Router = mux.NewRouter()

	app.setRouters()
}

func (app *App) setRouters() {
	app.Router.HandleFunc("/products", app.getAllAddress).Methods("get")
}

func (app *App) getAllAddress(w http.ResponseWriter, r *http.Request) {
	handlers.GetAllAddress(app.DB, w, r)
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
