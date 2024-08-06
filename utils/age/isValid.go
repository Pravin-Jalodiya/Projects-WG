package age

import "projects/config"

func ValidAge(age int) bool {
	if age < config.MIN_AGE || age > config.MAX_AGE {
		return false
	}
	return true
}
