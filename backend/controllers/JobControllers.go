package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	db "github.com/its-dev24/off-campus-placement-tracker/DB"
)

//Handler For retriving All job Listing

func HandleReadAll(w http.ResponseWriter, r *http.Request) {
	jobs, err := db.ReadAllApplications()
	if err != nil {
		json.NewEncoder(w).Encode("Error while reading jobs" + err.Error())
		log.Fatal("Error while reading jobs")
		return
	}
	if jobs == nil {
		json.NewEncoder(w).Encode("No jobs to display")
		return
	}
	json.NewEncoder(w).Encode(jobs)

}

//Handler For Deleting a Single Job

func HandleDeleteOne(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	deletedCount, err := db.DeleteASingleApplication(id)
	if err != nil {
		json.NewEncoder(w).Encode("Error While Deleteing A Single job : " + err.Error())
		log.Fatal("Error While Deleteing A Single job : ", err)
		return
	}
	json.NewEncoder(w).Encode("No of Items Deleted : " + strconv.Itoa(deletedCount))
}

//Function To Delete All Jobs

func HandeleDeleteAll(w http.ResponseWriter, r *http.Request) {
	deleteCount, err := db.DeleteAllApplications()
	if err != nil {
		json.NewEncoder(w).Encode("Error While Deleteing All job : " + err.Error())
		log.Fatal("Error While Deleteing All job : ", err)
		return
	}
	json.NewEncoder(w).Encode("No of Items Deleted : " + strconv.Itoa(deleteCount))

}
