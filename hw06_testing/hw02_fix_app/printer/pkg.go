package printer

import (
	"fmt"

	"github.com/Mr-Cheen1/home_work/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	for _, employee := range staff {
		fmt.Println(employee)
	}
}
