package main

import (
	"os"
	"testing"

	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw06_testing/hw02/reader"
	"github.com/ZhdanovskikhAV/otus_home_work_basic/hw06_testing/hw02/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFixApp(t *testing.T) {
	err := FixApp()
	require.NoError(t, err)
}

func TestReadJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []types.Employee
	}{
		{
			name: "Correct JSON",
			input: `[
				{"userId": 1, "age": 30, "name": "Alice", "departmentId": 101},
				{"userId": 2, "age": 25, "name": "Bob", "departmentId": 102}
			]`,
			expected: []types.Employee{
				{UserID: 1, Age: 30, Name: "Alice", DepartmentID: 101},
				{UserID: 2, Age: 25, Name: "Bob", DepartmentID: 102},
			},
		},
		{
			name: "Incorrect JSON",
			input: `[
				{"userId": 1, "age": 30, "name": "Alice", "departmentId": 101},
				{"userId": 2, "age": 25, "name": "Bob", "departmentId": 102
			]`, // недостающая закрывающая скобка
			expected: []types.Employee{},
		},
		{
			name:  "Non-JSON input",
			input: `This is not a JSON file.`,
			//expected: []types.Employee{},
			expected: []types.Employee{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempFile, err := os.CreateTemp("", "test_data.json")
			require.NoError(t, err)
			defer os.Remove(tempFile.Name())

			_, err = tempFile.WriteString(tt.input)
			require.NoError(t, err)

			err = tempFile.Close()
			require.NoError(t, err)

			employees, err := reader.ReadJSON(tempFile.Name())

			if len(tt.expected) > 0 {
				require.NoError(t, err, "Expected no error but got: %v", err)
				assert.Equal(t, tt.expected, employees, "Expected employee slice to match")
			} else {
				require.Error(t, err, "Expected an error but got none")
				assert.Len(t, employees, 0, "Expected no employees to be returned")
			}
		})
	}
}
