package main

import "testing"

func TestStubDatabase(t *testing.T) {
	// Initialize stub database
	stubDB := &StubDatabase{
		data: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	// Verify result for "key1"
	result, err := stubDB.GetData("key1")
	if err != nil || result != "value1" {
		t.Errorf("Expected 'value1', got '%s', error: %v", result, err)
	}

	// Verify result for a non-existent key
	result, err = stubDB.GetData("key3")
	if err == nil || result != "" {
		t.Errorf("Expected error for 'key3', got '%s', error: %v", result, err)
	}
}
