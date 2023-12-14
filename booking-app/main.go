package main

import (
	"fmt"
	// "os/exec"
	// "strings"
	"time"
)

type userData struct {
	firstName          string
	lastName           string
	email              string
	numberOfTickets    uint8
	hasAllowedFeedback bool
}

func main() {
	//Global variables
	var conferenceName string = "Go reference"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50
	// var bookings []string

	var users = make([]userData, 0)
	users = append(users, userData{
		firstName: "Gabby",
		lastName: "WATCHO",
		email: "",
		numberOfTickets: 10,
		hasAllowedFeedback: false,
	})

	fmt.Println(users)

	//Init messages
	greetUsers(conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickects and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	// for {
	// 	//Get user input
	// 	var userFirstName string
	// 	var userLastName string
	// 	var userEmail string
	// 	var userTickets uint8

	// 	fmt.Print("Enter your first name: ")
	// 	fmt.Scan(&userFirstName)

	// 	fmt.Print("Enter your last name: ")
	// 	fmt.Scan(&userLastName)

	// 	fmt.Print("Enter your email: ")
	// 	fmt.Scan(&userEmail)

	// 	//logic for booking
	// 	fmt.Print("Enter your number of tickets: ")
	// 	fmt.Scan(&userTickets)

	// 	for userTickets > remainingTickets {
	// 		//logic for booking
	// 		fmt.Printf("The is only %v tickets left\n", remainingTickets)
	// 		fmt.Print("Enter your number of tickets: ")
	// 		fmt.Scan(&userTickets)
	// 	}
	// 	remainingTickets = remainingTickets - userTickets
	// 	bookings = append(bookings, userFirstName+" "+userLastName)

	// 	firstNames := []string{}
	// 	for _, booking := range bookings {
	// 		var name = strings.Fields(booking)
	// 		firstNames = append(firstNames, name[0])
	// 	}

	// 	fmt.Printf("First name: %v \n", firstNames)
	// 	exec.Command("clear")

	// 	fmt.Printf("Thank you for booking %v tickets.\n", userTickets)
	// 	fmt.Printf("There is %v tickets remaining.\n", remainingTickets)
	// }
}

func greetUsers(confName string) {
	fmt.Println("Welcome to", confName, "booking application")
}
