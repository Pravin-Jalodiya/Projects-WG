package age

import "projects/config"

func VerifyAge(age int) bool {
	if age < config.VALID_AGE {
		return false
	} else {
		return true
	}
}
