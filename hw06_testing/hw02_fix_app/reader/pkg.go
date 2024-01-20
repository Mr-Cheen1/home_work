package reader

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Mr-Cheen1/home_work/hw02_fix_app/types"
)

var DefaultReader = ReadJSON

func ReadJSON(filePath string) ([]types.Employee, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var employees []types.Employee
	err = json.Unmarshal(file, &employees)
	if err != nil {
		return nil, err
	}

	return employees, nil
}
