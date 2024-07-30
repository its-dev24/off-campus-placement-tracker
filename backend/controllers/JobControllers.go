package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	db "github.com/its-dev24/off-campus-placement-tracker/DB"
	"github.com/its-dev24/off-campus-placement-tracker/modal"
)

//Handler For retriving All job Listing

func HandleReadAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	deletedCount, err := db.DeleteASingleApplication(id)
	if err != nil {
		json.NewEncoder(w).Encode("Error While Deleteing A Single job : " + err.Error())
		log.Fatal("Error While Deleteing A Single job : ", err)
		return
	}
	json.NewEncoder(w).Encode("No of Items Deleted : " + strconv.Itoa(deletedCount))
}

//Handler For Delete All Jobs

func HandeleDeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deleteCount, err := db.DeleteAllApplications()
	if err != nil {
		json.NewEncoder(w).Encode("Error While Deleteing All job : " + err.Error())
		log.Fatal("Error While Deleteing All job : ", err)
		return
	}
	json.NewEncoder(w).Encode("No of Items Deleted : " + strconv.Itoa(deleteCount))

}

//Handler For update Jobs

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	if r.Body == http.NoBody {
		json.NewEncoder(w).Encode("Request Body is Empty")
		return
	}
	var jobBody modal.Job
	err := json.NewDecoder(r.Body).Decode(&jobBody)
	if err != nil {
		log.Fatal("Error while Parsing Body : jobController.go : ", err)
	}
	if jobBody.IsEmpty() {
		json.NewEncoder(w).Encode("Json Body Is Empty Please Fill All The Fields")
		return
	}
	updateCount, err := db.UpdateJobs(id, jobBody)
	if err != nil {
		json.NewEncoder(w).Encode("Error While Updating Job..")
		log.Fatal("Error while updating Job : JobController.go : ", err)
	}
	json.NewEncoder(w).Encode("No of Fileds Updated " + strconv.Itoa(updateCount))
}

//Handler For Inserting Job

func HandleInsert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == http.NoBody {
		json.NewEncoder(w).Encode("The Request Body Is Empty")
		return
	}
	var job modal.Job
	json.NewDecoder(r.Body).Decode(&job)
	if job.IsEmpty() {
		json.NewEncoder(w).Encode("Json Body is Empty Fill All The Fields...")
		return
	}
	insertId, err := db.InsertJobs(job)
	if err != nil {
		json.NewEncoder(w).Encode("Error While Inserting Jobs")
		log.Fatal("Error While Inserting job : jobcontroller.go : ", err)
	}
	json.NewEncoder(w).Encode("ID of new Job is : " + insertId)
}
