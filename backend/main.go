package main

import (
	"fmt"

	db "github.com/its-dev24/off-campus-placement-tracker/DB"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error while Loading Environment Variable : main.go : ", err)
	// }
	// port := os.Getenv("PORT")
	// mux := router.Router()

	// server := http.Server{
	// 	Addr:    port,
	// 	Handler: mux,
	// }

	// err = server.ListenAndServe()
	// if err != nil {
	// 	log.Fatal("Error While Running Server : main.go : ", err)
	// }
	fmt.Println(db.MongoCol)
}
