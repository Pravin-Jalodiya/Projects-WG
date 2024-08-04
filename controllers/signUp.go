package controllers

import (
	"bufio"
	"fmt"
	"os"
	"projects/config"
	"projects/models"
	"projects/utils"
	"strconv"
	"strings"
)

var (
	reader   = bufio.NewReader(os.Stdin)
	age      int
	username string
	password string
)

func SignUp() {

	for {
		fmt.Print("Enter your desired username: ")
		user, err := reader.ReadString('\n')
		user = strings.TrimSuffix(user, "\n")
		user = strings.TrimSpace(user)
		if err != nil {
			fmt.Println("Error reading input.")
			//log errors
			return
		} else {
			if !utils.UniqueUser(user) {
				fmt.Println("Username already exists. Please choose a different username.")
			} else {
				username = user
				break
			}
		}
	}

	for {
		fmt.Print("Enter your password: ")
		pass, err := reader.ReadString('\n')
		pass = strings.TrimSuffix(pass, "\n")
		pass = strings.TrimSpace(pass)
		if err != nil {
			fmt.Println("Error reading input.")
			return
		} else {
			if !utils.PasswordValidator(pass) {
				fmt.Println("Weak password. Try another password (password should be at least 8  characters long and must have at least 1 lowercase, 1 uppercase, 1 special and 1 digit characters.")
			} else {
				password = pass
				break
			}
		}
	}

	for {
		fmt.Print("Enter your age: ")
		ageString, err := reader.ReadString('\n')
		ageString = strings.TrimSuffix(ageString, "\n")
		ageString = strings.TrimSpace(ageString)
		if err != nil {
			fmt.Println("Error reading age.")
			return
		} else {
			age, err = strconv.Atoi(ageString)
			if err != nil {
				fmt.Println("Not a number")
				continue
			}
			if utils.ValidAge(age) {
				break
			} else {
				fmt.Println("Please enter a valid age (range 1-150)")
			}
		}
	}

	if !utils.VerifyAge(age) {
		fmt.Println("Sign up failed : Minimum age criteria not met.")
		return
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password.")
		return
	}

	newUser := models.UserData{Username: username, Password: hashedPassword, ToDo: utils.Courses}

	ok, err := utils.FWriterUser(config.USER_FILE, newUser)
	if ok {
		fmt.Println("Sign up successful!")
	} else {
		fmt.Println("Sign up failed : ", err)
	}
}
