package services

import (
	"fmt"
	"projects/middleware"
)

func Logout() {
	middleware.ActiveUser = ""
	fmt.Println("User Logged out")
	return
}
