package types

import "fmt"

type Employee struct {
	UserID       int
	Age          int
	Name         string
	DepartmentID int
}

func (e Employee) String() string {
	return fmt.Sprintf(
		"User ID: %d; Age: %d; Name: %s; Department ID: %d",
		e.UserID, e.Age, e.Name, e.DepartmentID,
	)
}
