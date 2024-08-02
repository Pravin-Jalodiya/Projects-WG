package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("\nPlease select an option:")
		fmt.Println("1. Sign Up")
		fmt.Println("2. Log In")
		fmt.Println("3. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			signUp()
		case 2:
			login()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}

func signUp() {
	var username, password string
	var age int

	fmt.Print("Enter your desired username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	if age < 12 {
		fmt.Println("Sign up unsuccessful: Minimum age criteria not met.")
	} else {
		registerUser(username, password)
	}
}

func existingUser(username string) bool {

	file, err := os.OpenFile("users.txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	usernames := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		parts := strings.SplitN(line, "\t", 2)
		if len(parts) == 2 {
			//fmt.Println(parts[0])
			usernames[parts[0]] = true
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
		return false
	}
	username = fmt.Sprintf("username: %s", username)
	if _, exists := usernames[username]; exists {
		fmt.Println("Username already exists. Please choose a different username.")
		return false
	}
	return true
}
func registerUser(username, password string) {
	//file, err := os.Open("users.txt")
	if existingUser(username) {
		file, err := os.OpenFile("users.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("username: %s\tpassword: %s\n", username, password))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println("Sign up successful!")
	} else {
		return
	}
}

func login() {
	var username, password string

	fmt.Print("Enter your username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	userExists := false
	file, err := os.OpenFile("users.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "\t", 2)
		if len(parts) == 2 && parts[0] == fmt.Sprintf("username: %s", username) && parts[1] == fmt.Sprintf("password: %s", password) {
			userExists = true
			break
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading file:", scanner.Err())
		return
	}

	if userExists {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Login failed: Invalid username or password")
	}

}
