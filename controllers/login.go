package controllers

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
	"projects/middleware"
	"projects/services"
	password2 "projects/utils/password"
	"projects/utils/readers"
	"strings"
	"syscall"
)

func Login() {

	successEmoji := "✅"
	green := color.New(color.FgGreen).SprintFunc()

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
	secret, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	fmt.Println()
	pass := string(secret)
	pass = strings.TrimSpace(pass)
	password = pass

	if storedHash, exists := readers.UserMap[username]; exists {
		if password2.VerifyPassword(password, storedHash) {
			fmt.Println(green(successEmoji), green("Log In successful!"))
			middleware.Auth(username)
			services.Main()
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
