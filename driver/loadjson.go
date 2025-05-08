package driver

import (
	"encoding/json"
	"os"

	"github.com/aponx/book-management/app/domain"
)

// NewLoadJson return array object of book
func NewLoadJson(file string) (*[]domain.Book, error) {
	// Read the JSON file
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a map
	var data []domain.Book
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

func UpdateJsonFil(books []domain.Book, file string) error {
	// Marshal the updated data back to JSON
	updatedData, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}

	// Write the updated JSON data to the file
	err = os.WriteFile(file, updatedData, 0644)
	if err != nil {
		return err
	}

	return nil
}
