package helper

import "strings"

func ValidateUserInput(first_Name string, last_Name string, email string, user_Ticket uint, remaininng_Ticket uint) (bool, bool, bool) {
	isValidName := len(first_Name) >= 2 && len(last_Name) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := user_Ticket > 0 && user_Ticket <= remaininng_Ticket
	return isValidName, isValidEmail, isValidTicket
}