package reader

import (
	"encoding/json"
	"fmt"
	"hw02_fix_app/types"
	"io"
	"os"
)

func ReadJSON(filePath string, _ int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	_byte, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	var data []types.Employee

	err = json.Unmarshal(_byte, &data)

	res := data

	return res, err
}
