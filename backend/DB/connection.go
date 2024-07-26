package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCol *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading Environment Variable : connection.go : ", err)
	}
	connectionString := os.Getenv("MONGO_CLIENT")
	database := os.Getenv("DB_NAME")
	colName := os.Getenv("COL_NAME")
	clientOptions := options.Client().ApplyURI(connectionString)
	Client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error while Creating Mongo Client : connection.go : ", err)
	}
	MongoCol = Client.Database(database).Collection(colName)
}
