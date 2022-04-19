package main

import (
	"BOOKING-APP/helper"
	"fmt"
	"sync"
	"time"
)

var conference_Name = "Go Conference"
const conference_Ticket = 50
var remaininng_Ticket uint = 50
var bookings = make([]UserData, 0)


type UserData struct {
	first_Name string
	last_Name string
	email string
	Number_of_Ticket uint
}


var wg = sync.WaitGroup{}

func main(){
	

	Greet_User()
	

	for {
		first_Name, last_Name, email, user_Ticket := GetUserInput()
		isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(first_Name, last_Name, email, user_Ticket, remaininng_Ticket)

		if isValidName && isValidEmail && isValidTicket {
			bookTicket(user_Ticket, first_Name, last_Name, email)

			wg.Add(1)
			go sendTicket(user_Ticket, first_Name, last_Name, email)

			first_Names := GetFirstNames()
			GetFirstNames()

			fmt.Printf("The first names of bookings are : %v\n", first_Names)
			if remaininng_Ticket == 0{
			//end of program
				fmt.Println("The conference is booked out, come backnext year.")
				break	
			}
		}else{
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contains @")
			}
			if !isValidTicket {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Println("Your input data is invalid, try again !")
		}
	}
	wg.Wait()
}

func Greet_User() {
	fmt.Printf("Welcome to our %v conference !", conference_Name)
	fmt.Println("Welcome to"  + conference_Name + "booking application!")
	fmt.Println("We have total a of ", conference_Ticket, " tickets and ", remaininng_Ticket, " are still available")
	fmt.Println("Get your tickets here to attend")
}

func GetFirstNames() []string {
	first_Names := []string{}
	for _, booking := range bookings {
		first_Names = append(first_Names, booking.first_Name)
	}
	return first_Names
}




func GetUserInput() (string, string, string, uint) {
	var first_Name string
	var last_Name string
	var email string
	var user_Ticket uint

	fmt.Println("Enter your first name :")
	fmt.Scan(&first_Name)

	fmt.Println("Enter your last name :")
	fmt.Scan(&last_Name)

	fmt.Println("Enter your email :")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets :")
	fmt.Scan(&user_Ticket)

	return first_Name, last_Name, email, user_Ticket
}

func bookTicket(user_Ticket uint, first_Name string, last_Name string, email string) {
	remaininng_Ticket = remaininng_Ticket - user_Ticket

	var userData = UserData {
		first_Name: first_Name,
		last_Name: last_Name,
		email: email,
		Number_of_Ticket: user_Ticket,		
	}


	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets! You'll receive a confirmation email at %v\n", first_Name, last_Name, user_Ticket, email)
	fmt.Printf("%v tickets remaining for %v\n", remaininng_Ticket, conference_Name)
}


func sendTicket(user_Ticket uint, first_Name string, last_Name string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", user_Ticket, first_Name, last_Name)
	fmt.Println("##################")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done()
}