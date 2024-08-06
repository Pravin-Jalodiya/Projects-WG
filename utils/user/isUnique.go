package user

import "projects/utils/readers"

func IsUnique(username string) bool {
	if _, exists := readers.UserMap[username]; exists {
		return false
	}
	return true
}
