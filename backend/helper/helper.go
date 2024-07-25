//TODO: Decide To Write a Func to input a slice of slices and return a slice of structs

package helper

import (
	"log"
	"strconv"

	"github.com/its-dev24/off-campus-placement-tracker/modal"
)

func ConvertSliceToStruct(job []interface{}) modal.Job {
	id, err := strconv.Atoi(job[0].(string))
	if err != nil {
		log.Fatal("Error while converting id : ", err)
	}
	resultJob := modal.Job{
		Id:      int(id),
		Company: job[1].(string),
		JobRole: job[2].(string),
		Status:  job[3].(string),
		WebSite: job[4].(string),
	}
	return resultJob
}
