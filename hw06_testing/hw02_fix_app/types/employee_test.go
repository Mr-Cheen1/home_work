package types

import (
	"testing"
)

func TestEmployeeString(t *testing.T) {
	tests := []struct {
		emp  Employee
		want string
	}{
		{
			emp:  Employee{UserID: 1, Age: 30, Name: "Alice", DepartmentID: 5},
			want: "User ID: 1; Age: 30; Name: Alice; Department ID: 5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.emp.Name, func(t *testing.T) {
			got := tt.emp.String()
			if got != tt.want {
				t.Errorf("Employee.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
