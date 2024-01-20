package main

import (
	"fmt"
	"os"

	"github.com/Mr-Cheen1/home_work/hw02_fix_app/printer"
	"github.com/Mr-Cheen1/home_work/hw02_fix_app/reader"
	"github.com/Mr-Cheen1/home_work/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}

	err := ProcessDataFile(path)
	if err != nil {
		fmt.Printf("Error processing data file: %v\n", err)
		return
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening data file: %v\n", err)
		return
	}
	defer file.Close()

	staff, err := reader.ReadJSON(file)
	if err != nil {
		fmt.Printf("Error reading data file: %v\n", err)
		return
	}
	printer.PrintStaff(staff)
}

func ProcessDataFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	employees, err := reader.ReadJSON(file)
	if err != nil {
		return err
	}
	for i := range employees {
		employees[i] = ProcessEmployee(employees[i])
	}
	return nil
}

func ProcessEmployee(employee types.Employee) types.Employee {
	if employee.DepartmentID == 2 {
		employee.Age += 1
	}
	return employee
}
