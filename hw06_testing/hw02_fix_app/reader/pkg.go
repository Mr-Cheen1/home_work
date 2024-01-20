package reader

import (
	"encoding/json"
	"io"

	"github.com/Mr-Cheen1/home_work/hw02_fix_app/types"
)

func ReadJSON(r io.Reader) ([]types.Employee, error) {
	var employees []types.Employee
	err := json.NewDecoder(r).Decode(&employees)
	if err != nil {
		return nil, err
	}
	return employees, nil
}
