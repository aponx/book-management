package driver

import (
	"encoding/json"
	"os"

	"github.com/aponx/book-management/app/domain"
)

// NewPostgreDatabase return gorp dbmap object with postgre options param
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
