package course

import (
	"projects/models"
	"projects/utils/readers"
	"testing"
)

func TestCidFound(t *testing.T) {
	// Reset UserStore before each test to avoid state leakage between tests
	defer resetUserStore()
	// Test case 1: CID is found for the current user
	user1 := models.UserData{Username: "user1"}
	toDo1 := models.Course{CID: 123}
	user1.ToDo = append(user1.ToDo, toDo1)
	readers.UserStore = append(readers.UserStore, user1)

	if !cidFound(123, "user1") {
		t.Errorf("Expected CID 123 to be found for user1, but it was not")
	}

	// Test case 2: CID is not found for the current user
	if cidFound(456, "user1") {
		t.Errorf("Expected CID 456 not to be found for user1, but it was")
	}

	// Test case 3: The user store is empty
	resetUserStore() // Ensure UserStore is empty
	if cidFound(123, "nonexistent_user") {
		t.Errorf("Expected no CIDs to be found in an empty UserStore, but a CID was reported")
	}
}

// Helper function to reset UserStore for isolation between tests
func resetUserStore() {
	readers.UserStore = []models.UserData{} // Assuming UserStore is a slice of models.User
}
