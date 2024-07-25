package sheetshelper

import (
	"context"
	"log"
	"os"

	"github.com/its-dev24/off-campus-placement-tracker/helper"
	"github.com/its-dev24/off-campus-placement-tracker/modal"
	"github.com/joho/godotenv"
)

var sheetsID string

func init() {
	var err = godotenv.Load()
	if err != nil {
		log.Fatal("Error While Loading Env : sheetsHelper : ", err)
	}
	sheetsID = os.Getenv("SHEET_ID")
}

// Function To read all Job Applications
func ReadApplications() []modal.Job {
	result, err := Client.Spreadsheets.Values.Get(sheetsID, "Sheet1!A:E").Context(context.Background()).Do()
	if err != nil {
		log.Fatal("Error While retriving Data : ", err)
	}
	//FIXME: Need to fix this Function
	var jobs []modal.Job
	for _, val := range result.Values {
		job := helper.ConvertSliceToStruct(val)
		jobs = append(jobs, job)
	}
	return jobs
}

//TODO:Function To Insert A Job Application

//TODO: Function To Update A Job Application

//TODO: Function To Delete A Job Application
