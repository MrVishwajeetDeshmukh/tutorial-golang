package main

import "errors"

// MockAPIClient is a fake object for testing.
type MockAPIClient struct{}

// FetchData returns fixed data.
func (m *MockAPIClient) FetchData(endpoint string) (string, error) {
	if endpoint == "test/valid" {
		return "Mock Data", nil
	}
	return "", errors.New("invalid endpoint")
}
