package db

import (
	"context"

	"github.com/its-dev24/off-campus-placement-tracker/modal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	jobId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": jobId}
	deleteResult, err := MongoCol.DeleteOne(context.Background(), filter)
	if err != nil {
		return 0, nil
	}
	return int(deleteResult.DeletedCount), nil
}

//Function To Update VAlues

func UpdateJobs(id string, jobs modal.Job) (int, error) {
	jobId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": jobId}
	update := bson.M{"$set": bson.M{"company": jobs.Company, "jobrole": jobs.JobRole, "status": jobs.Status, "website": jobs.WebSite}}
	updateResult, err := MongoCol.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, nil
	}
	return int(updateResult.ModifiedCount), nil
}

//Function To Insert Values

func InsertJobs(job modal.Job) (string, error) {
	insertResult, err := MongoCol.InsertOne(context.Background(), job)
	if err != nil {
		return "", err
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}
