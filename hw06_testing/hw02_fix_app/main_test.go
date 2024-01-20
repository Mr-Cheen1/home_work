package main

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Mr-Cheen1/home_work/hw02_fix_app/reader"
	"github.com/Mr-Cheen1/home_work/hw02_fix_app/types"
)

func TestProcessDataFile(t *testing.T) {
	originalReader := reader.DefaultReader
	defer func() { reader.DefaultReader = originalReader }()

	tests := []struct {
		name    string
		path    string
		mock    func(string) ([]types.Employee, error)
		want    []types.Employee
		wantErr bool
	}{
		{
			name: "correct data",
			path: "data.json",
			mock: func(string) ([]types.Employee, error) {
				return []types.Employee{
					{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3},
					{UserID: 11, Age: 30, Name: "George", DepartmentID: 2},
				}, nil
			},
			want: []types.Employee{
				{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3},
				{UserID: 11, Age: 30, Name: "George", DepartmentID: 2},
			},
			wantErr: false,
		},
		{
			name: "file not found",
			path: "nonexistent.json",
			mock: func(string) ([]types.Employee, error) {
				return nil, errors.New("file not found")
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader.DefaultReader = tt.mock
			err := ProcessDataFile(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessDataFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProcessEmployee(t *testing.T) {
	tests := []struct {
		name string
		in   types.Employee
		want types.Employee
	}{
		{
			name: "age increment for department 2",
			in:   types.Employee{UserID: 11, Age: 30, Name: "George", DepartmentID: 2},
			want: types.Employee{UserID: 11, Age: 31, Name: "George", DepartmentID: 2},
		},
		{
			name: "no change for other departments",
			in:   types.Employee{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3},
			want: types.Employee{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProcessEmployee(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessEmployee() received = %v, pending %v", got, tt.want)
			}
		})
	}
}
