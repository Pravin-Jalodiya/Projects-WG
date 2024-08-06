package writers

import (
	"encoding/json"
	"log"
	"os"
	"projects/models"
	"projects/utils/readers"
)

func FWriterToDo(f string, newUser []models.UserData) (bool, error) {

	users := readers.FReaderUser(f, os.O_CREATE|os.O_APPEND|os.O_RDWR)

	users = newUser

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
	readers.SyncUserData()
	return true, nil
}
