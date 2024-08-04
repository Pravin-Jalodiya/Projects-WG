package utils

func UniqueUser(username string) bool {
	if _, exists := UserMap[username]; exists {
		return false
	}
	return true
}
