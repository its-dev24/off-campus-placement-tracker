package modal

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Company string             `json:"company,omitempty" bson:"company,omitempty"`
	JobRole string             `json:"jobrole,omitempty" bson:"jobrole,omitempty"`
	Status  string             `json:"status,omitempty" bson:"status,omitempty"`
	WebSite string             `json:"wesite,omitempty" bson:"wesite,omitempty"`
}

func (j *Job) IsEmpty() bool {
	return j.Company == "" || j.JobRole == "" || j.Status == "" || j.WebSite == ""
}
