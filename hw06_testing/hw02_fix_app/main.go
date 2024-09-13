package main

import (
	"fmt"

	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw06_testing/hw02_fix_app/printer"
	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw06_testing/hw02_fix_app/reader"
	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw06_testing/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	var err error
	fmt.Scanln(&path)

	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path, -1)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
