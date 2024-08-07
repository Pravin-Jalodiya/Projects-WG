package controllers

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"projects/LLM"
	"projects/config"
	"projects/models"
	age2 "projects/utils/age"
	password2 "projects/utils/password"
	user2 "projects/utils/user"
	"projects/utils/writers"
	"strconv"
	"strings"
	"syscall"
)

var (
	reader   = bufio.NewReader(os.Stdin)
	age      int
	username string
	password string
)

func SignUp() {

	for {
		fmt.Print("\nEnter your desired username: ")
		user, err := reader.ReadString('\n')
		user = strings.TrimSuffix(user, "\n")
		user = strings.TrimSpace(user)
		if err != nil {
			fmt.Println("Error reading input.")
			//log errors
			return
		} else {
			if !user2.IsUnique(user) {
				fmt.Println("Username already exists. Please choose a different username.")
				fmt.Println("Need help with finding a unique username?(y/n)")
				choice, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading input.")
				}
				choice = strings.TrimSuffix(choice, "\n")
				choice = strings.TrimSpace(choice)
				if choice == "y" {
					suggestedUsernames := LLM.UsernameSuggestion(user)
					fmt.Printf("Username suggestions\n%s", suggestedUsernames)
				}
			} else {
				username = user
				break
			}
		}
	}

	for {
		fmt.Print("Enter your password: ")
		secret1, err := term.ReadPassword(syscall.Stdin)
		fmt.Println()
		pass1 := string(secret1)
		pass1 = strings.TrimSpace(pass1)
		if err != nil {
			fmt.Println("Error reading input.")
			return
		} else {
			if !password2.ValidatePass(pass1) {
				fmt.Println("Weak password. Try another password (should be at least 8  characters long and must have at least 1 lowercase, 1 uppercase, 1 special and 1 digit characters.)")
				fmt.Println("Need help with finding the right password?(y/n)")
				for {
					choice, err := reader.ReadString('\n')
					if err != nil {
						fmt.Println("Error reading input.")
					}
					choice = strings.TrimSuffix(choice, "\n")
					choice = strings.TrimSpace(choice)
					if choice == "y" {
						suggestedPass := LLM.PasswordSuggestion()
						fmt.Println("Password suggestion : " + suggestedPass)
						break
					} else {
						if choice == "n" {
							break
						}
						fmt.Println("Invalid input. Please try again.")
					}
				}
			} else {
				fmt.Print("Enter your password again: ")
				secret2, err := term.ReadPassword(syscall.Stdin)
				fmt.Println()
				if err != nil {
					fmt.Println("Error reading input.")
					return
				}
				pass2 := string(secret2)
				pass2 = strings.TrimSpace(pass2)
				if pass1 == pass2 {
					password = pass2
					break
				}
				fmt.Println("Passwords don't match. Try again.")
			}
		}
	}

	for {
		fmt.Print("Enter your age: ")
		ageString, err := reader.ReadString('\n')
		fmt.Println()
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
			if age2.ValidAge(age) {
				break
			} else {
				fmt.Println("Please enter a valid age (range 1-150)")
			}
		}
	}

	if !age2.VerifyAge(age) {
		fmt.Println("Sign up failed : Minimum age criteria not met.")
		return
	}

	hashedPassword, err := password2.HashPass(password)
	if err != nil {
		fmt.Println("Error hashing password.")
		return
	}

	emptyCourses := []models.Course{}
	emptyDailyStatus := []models.DailyStatus{}
	newUser := models.UserData{Username: username, Password: hashedPassword, ToDo: emptyCourses, DailyStatus: emptyDailyStatus}

	ok, err := writers.FWriterUser(config.USER_FILE, newUser)
	if ok {
		fmt.Println("Sign up successful!")
	} else {
		fmt.Println("Sign up failed : ", err)
	}
}
