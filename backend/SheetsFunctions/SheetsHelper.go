package sheetshelper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

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
func ReadApplications() {
	result, err := Client.Spreadsheets.Values.Get(sheetsID, "Sheet1!A:D").Context(context.Background()).Do()
	if err != nil {
		log.Fatal("Error While retriving Data : ", err)
	}
	job, _ := json.Marshal(result.Values)
	jobReader := bytes.NewReader(job)
	//FIXME: Need to fix this Function
	var jobs []modal.Job
	json.NewDecoder(jobReader).Decode(&jobs)
	fmt.Println(jobs)

}

//TODO:Function To Insert A Job Application

//TODO: Function To Update A Job Application

//TODO: Function To Delete A Job Application
