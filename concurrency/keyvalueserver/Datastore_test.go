package main

import "testing"

func TestDataStoreMethods(t *testing.T) {
	// Initialize DataStore
	ds := DataStore{storage: make(map[string]int64)}

	// Test AddItem method
	ds.AddItem("key1", 10)
	if ds.storage["key1"] != 10 {
		t.Errorf("AddItem failed: expected value 10, got %d", ds.storage["key1"])
	}

	// Test IncItem method with existing key
	success, err := ds.IncItem("key1")
	if !success {
		t.Errorf("IncItem failed: %v", err)
	}
	if ds.storage["key1"] != 11 {
		t.Errorf("IncItem failed: expected value 11, got %d", ds.storage["key1"])
	}

	// Test IncItem method with non-existing key
	success, err = ds.IncItem("nonexistent")
	if success {
		t.Errorf("IncItem failed: expected failure, but got success")
	}
	expectedError := "Cannot increment. THE KEY DOES NOT EXIST!!!!"
	if err.Error() != expectedError {
		t.Errorf("IncItem failed: expected error '%s', got '%s'", expectedError, err.Error())
	}

	// Test MultiplyItem method with existing key
	success, err = ds.MultiplyItem("key1", 2)
	if !success {
		t.Errorf("MultiplyItem failed: %v", err)
	}
	if ds.storage["key1"] != 22 {
		t.Errorf("MultiplyItem failed: expected value 22, got %d", ds.storage["key1"])
	}

	// Test MultiplyItem method with non-existing key
	success, err = ds.MultiplyItem("nonexistent", 2)
	if success {
		t.Errorf("MultiplyItem failed: expected failure, but got success")
	}
	expectedError = "Cannot mulitply. THE KEY DOES NOT EXIST!!!!"
	if err.Error() != expectedError {
		t.Errorf("MultiplyItem failed: expected error '%s', got '%s'", expectedError, err.Error())
	}

	// Test GetItem method with existing key
	ok, value := ds.GetItem("key1")
	if !ok {
		t.Errorf("GetItem failed: key 'key1' not found")
	}
	if value != 22 {
		t.Errorf("GetItem failed: expected value 22, got %d", value)
	}

	// Test GetItem method with non-existing key
	ok, _ = ds.GetItem("nonexistent")
	if ok {
		t.Errorf("GetItem failed: key 'nonexistent' should not exist")
	}
}
