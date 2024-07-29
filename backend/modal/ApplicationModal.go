package modal

type Job struct {
	Id      int
	Company string `json:"company,omitempty" bson:"company,omitempty"`
	JobRole string `json:"jobrole,omitempty" bson:"jobrole,omitempty"`
	Status  string `json:"status,omitempty" bson:"status,omitempty"`
	WebSite string `json:"wesite,omitempty" bson:"wesite,omitempty"`
}

func (j *Job) IsEmpty() bool {
	return j.Company == "" || j.JobRole == "" || j.Status == "" || j.WebSite == ""
}
