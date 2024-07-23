package authentication

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var Client *sheets.Service

func init() {

	godotenv.Load()
	credentialFile := os.Getenv("PATH_TO_GKEYS")
	var err error
	Client, err = sheets.NewService(context.Background(), option.WithCredentialsFile(credentialFile))
	if err != nil {
		log.Fatal("Could not create sheets client : ", err)
	}
}
