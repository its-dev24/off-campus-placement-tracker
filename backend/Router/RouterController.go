package router

import (
	"net/http"

	"github.com/its-dev24/off-campus-placement-tracker/controllers"
)

func Router() *http.ServeMux {
	var mux http.ServeMux
	mux.HandleFunc("GET /api/jobs", controllers.HandleReadAll)
	return &mux

}
