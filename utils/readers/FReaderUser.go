package readers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"projects/models"
)

var (
	UserStore []models.UserData
	UserMap   = map[string]string{}
)

func init() {
	SyncUserData()
	for _, v := range UserStore {
		UserMap[v.Username] = v.Password
	}
}

func FReaderUser(f string, flag int) []models.UserData {

	var users []models.UserData

	file, err := os.OpenFile(f, flag, 0644)
	if err != nil {
		fmt.Println("Error opening file")
		//log error
	}
	byteValue, _ := io.ReadAll(file)
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return nil
	}
	return users
}
