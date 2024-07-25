package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	sheetshelper "github.com/its-dev24/off-campus-placement-tracker/SheetsFunctions"
)

func handleReadAll(w http.ResponseWriter, r *http.Request) {
	jobs := sheetshelper.ReadApplications()
	err := json.NewEncoder(w).Encode(jobs)
	if err != nil {
		log.Fatal("Error While Reading all jobs : jobControllers :  ", err)
	}
}
