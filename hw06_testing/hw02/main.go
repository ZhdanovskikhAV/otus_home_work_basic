package main

import (
	"fmt"

	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw06_testing/hw02/printer"
	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw06_testing/hw02/reader"
)

func main() {
	FixApp()
}

func FixApp() error {
	var path string

	fmt.Print("Enter data file path (press Enter for default 'data.json'): ")
	fmt.Scanln(&path)

	if path == "" {
		path = "data.json"
	}

	staff, err := reader.ReadJSON(path)
	if err != nil {
		return fmt.Errorf("failed to read employees data: %w", err)
	}

	printer.PrintStaff(staff)
	return nil
}
