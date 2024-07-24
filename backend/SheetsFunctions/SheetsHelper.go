package sheetshelper

import (
	"context"
	"fmt"
	"log"
	"os"

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
	fmt.Println(result)

}

//TODO:Function To Insert A Job Application

//TODO: Function To Update A Job Application

//TODO: Function To Delete A Job Application
