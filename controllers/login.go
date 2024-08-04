package controllers

import (
	"fmt"
	"projects/middleware"
	"projects/services"
	"projects/utils"
	"strings"
)

func Login() {

	fmt.Print("Enter your username: ")
	user, err2 := reader.ReadString('\n')
	user = strings.TrimSuffix(user, "\n")
	user = strings.TrimSpace(user)
	if err2 != nil {
		fmt.Println("Error reading input.")
		//log errors
		return
	} else {
		username = user
	}

	fmt.Print("Enter your password: ")
	pass, err2 := reader.ReadString('\n')
	pass = strings.TrimSuffix(pass, "\n")
	pass = strings.TrimSpace(pass)
	if err2 != nil {
		fmt.Println("Error reading input.")
		//log errors
		return
	} else {
		password = pass
	}

	if storedHash, exists := utils.UserMap[username]; exists {
		if utils.VerifyPassword(password, storedHash) {
			fmt.Println("Login successful")
			middleware.Auth(username)
			services.Services()
		} else {
			fmt.Println("Incorrect username or password")
		}
	} else {
		fmt.Println("User not found. Do you wish to sign up?(y/n)")
		signupChoice, err := reader.ReadString('\n')
		signupChoice = strings.TrimSuffix(signupChoice, "\n")
		signupChoice = strings.TrimSpace(signupChoice)
		if err != nil {
			fmt.Println("Error reading input.")
			return
		}
		if signupChoice == "y" {
			SignUp()
		} else {
			return
		}

	}

}
