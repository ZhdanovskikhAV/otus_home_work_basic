package reader

import (
	"encoding/json"
	"os"

	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw06_testing/hw02/types"
)

// ReadJSON читает JSON-данные из файла и возвращает срез структур Employee.
func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err

	}
	defer f.Close()

	var data []types.Employee

	if err := json.NewDecoder(f).Decode(&data); err != nil {
		return nil, err

	}
	return data, nil
}
