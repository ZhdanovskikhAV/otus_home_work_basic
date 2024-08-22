package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw02_fix_app/types"
)

func ReadJSON(filePath string, _ int) ([]types.Employee, error) {

	f, err := os.Open(filePath)

	if err != nil {

		fmt.Printf("Error: %v", err)

		return nil, err

	}

	_byte, err := io.ReadAll(f)

	if err != nil {

		fmt.Printf("Error: %v", err)

		return nil, err

	}

	var data []types.Employee

	err = json.Unmarshal(_byte, &data)

	if err != nil {

		fmt.Println("Ошибка чтения JSON-данных:", err)

	}

	return data, err

}

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw02_fix_app/types"
)

func ReadJSON(filePath string, _ int) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	_byte, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	var data []types.Employee

	err = json.Unmarshal(_byte, &data)
	if err != nil {
		fmt.Println("Ошибка чтения JSON-данных:", err)
	}
	return data, err
}
