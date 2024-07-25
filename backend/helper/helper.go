//TODO: Decide To Write a Func to input a slice of slices and return a slice of structs

package helper

import "github.com/its-dev24/off-campus-placement-tracker/modal"

func convertSliceToStruct(job []string) modal.Job {
	resultJob := modal.Job{
		Company: job[0],
		JobRole: job[1],
		Status:  job[2],
		WebSite: job[3],
	}
	return resultJob
}
