package db

import (
	"context"

	"github.com/its-dev24/off-campus-placement-tracker/modal"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadAllApplications() ([]modal.Job, error) {
	cur, err := MongoCol.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var jobs []modal.Job
	cur.All(context.Background(), &jobs)
	return jobs, nil
}

