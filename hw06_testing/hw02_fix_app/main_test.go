package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Mr-Cheen1/home_work/hw02_fix_app/reader"
	"github.com/Mr-Cheen1/home_work/hw02_fix_app/types"
)

func TestProcessDataFile(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []types.Employee
		wantErr bool
	}{
		{
			name: "correct data",
			input: `[
				{"UserID": 10, "Age": 25, "Name": "Rob", "DepartmentID": 3},
				{"UserID": 11, "Age": 30, "Name": "George", "DepartmentID": 2}
			]`,
			want: []types.Employee{
				{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3},
				{UserID: 11, Age: 30, Name: "George", DepartmentID: 2},
			},
			wantErr: false,
		},
		{
			name:    "invalid json",
			input:   `invalid json`,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			got, err := reader.ReadJSON(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
