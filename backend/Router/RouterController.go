package router

import (
	"net/http"

	"github.com/its-dev24/off-campus-placement-tracker/controllers"
)

func Router() *http.ServeMux {
	var mux http.ServeMux
	mux.HandleFunc("GET /api/jobs", controllers.HandleReadAll)
	mux.HandleFunc("DELETE /api/jobs/{id}", controllers.HandleDeleteOne)
	mux.HandleFunc("DELETE /api/jobs", controllers.HandeleDeleteAll)
	mux.HandleFunc("PUT /api/jobs/{id}", controllers.HandleUpdate)
	return &mux

}
