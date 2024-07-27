package db

import (
	"context"

	"github.com/its-dev24/off-campus-placement-tracker/modal"
	"go.mongodb.org/mongo-driver/bson"
)

//Function To Find All Applications

func ReadAllApplications() ([]modal.Job, error) {
	cur, err := MongoCol.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var jobs []modal.Job
	cur.All(context.Background(), &jobs)
	return jobs, nil
}

//Function To Delete All Applications

func DeleteAllApplications() (int, error) {
	deleteResult, err := MongoCol.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}
	return int(deleteResult.DeletedCount), nil
}

//Function To Delete A Single Job

func DeleteASingleApplication(id string) (int, error) {

	filter := bson.M{"_id": id}
	deleteResult, err := MongoCol.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, nil
	}
	return int(deleteResult.DeletedCount), nil
}
