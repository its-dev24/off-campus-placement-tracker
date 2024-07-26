package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	db "github.com/its-dev24/off-campus-placement-tracker/DB"
)

func HandleReadAll(w http.ResponseWriter, r *http.Request) {
	jobs, err := db.ReadAllApplications()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal("Error while reading jobs")
		return
	}
	if jobs == nil {
		json.NewEncoder(w).Encode("No jobs to display")
		return
	}
	json.NewEncoder(w).Encode(jobs)

}
