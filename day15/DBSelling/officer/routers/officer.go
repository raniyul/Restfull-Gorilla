package routers

import (
	"day15/DBSelling/officer/controllers"

	"github.com/gorilla/mux"
)

func setOfficerRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/officer", controllers.GetOfficer).Methods("GET")
	return router
}
