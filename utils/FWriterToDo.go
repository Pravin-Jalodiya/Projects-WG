package utils

import (
	"encoding/json"
	"log"
	"os"
	"projects/config"
	"projects/models"
)

func FWriterToDo(f string, newUser []models.UserData) (bool, error) {

	users := FReaderUser(f, os.O_CREATE|os.O_APPEND|os.O_RDWR)

	users = newUser

	// Write updated users back to file
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Printf("Error marshaling data: %v\n", err)
		return false, err
	}

	err = os.WriteFile("users.json", jsonData, 0644)
	if err != nil {
		log.Printf("Error writing to file: %v\n", err)
		return false, err
	}
	UserStore = FReaderUser(config.USER_FILE, os.O_RDONLY|os.O_CREATE)
	return true, nil
}
