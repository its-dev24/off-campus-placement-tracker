package main

import (
	"fmt"

	sheetshelper "github.com/its-dev24/off-campus-placement-tracker/SheetsFunctions"
)

func main() {
	fmt.Println(sheetshelper.Client)
	sheetshelper.ReadApplications()
}
