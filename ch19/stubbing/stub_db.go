package main

import "errors"

// StubDatabase is a simple database implementation for testing.
type StubDatabase struct {
	data map[string]string
}

// GetData returns fixed data.
func (db *StubDatabase) GetData(key string) (string, error) {
	if value, exists := db.data[key]; exists {
		return value, nil
	}
	return "", errors.New("key not found")
}
