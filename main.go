package main

import (
	"booking_app/helper"
	"fmt"
	"sync"
	"time"
)

const totalTickets = 100

var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   int
}

var wait = sync.WaitGroup{}

func main() {

	greet()

	for {
		//input from user
		firstName, lastName, email, userTickets := userInput()

		//validate input
		valName, valEmail, valNumber := helper.ValidAte(firstName, lastName, email, userTickets, remainingTickets)

		if valName && valEmail && valNumber {
			//book function
			book(firstName, lastName, email, userTickets)
			// call function frst name here
			firstWords := frstname()
			fmt.Printf("First names are %v\n", firstWords)
			//Send ticket to gmail after booking
			wait.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			if remainingTickets == 0 {
				fmt.Println("Oops!Booked out.")
				break
			}
		} else {
			if !valName {
				fmt.Println("Invalid first or last name : ")
			}
			if !valEmail {
				fmt.Println("Invald email")
			}
			if !valNumber {
				fmt.Println("Invalid ticket number")
			}
		}

	}
	wait.Wait()

}

func greet() {
	fmt.Printf("Welcome to our %v booking app\n", conferenceName)
	fmt.Println("We have total", totalTickets, "tickets.The leftover are", remainingTickets, "Grab asap!!!")
	fmt.Println("Get your tickets here for booking")

}

func frstname() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func userInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Println("Enter user first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter user last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter user email: ")
	fmt.Scan(&email)

	fmt.Println("Enter No.of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func book(firstName string, lastName string, email string, userTickets int) {
	remainingTickets = remainingTickets - userTickets
	var userDetails = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}
	// //create a map()
	// var userDetails = make(map[string]string)
	// userDetails["first name"] = firstName
	// userDetails["last name"] = lastName
	// userDetails["email"] = email
	// userDetails["user tickets"] = strconv.FormatInt(int64(userTickets), 10)
	bookings = append(bookings, userDetails)

	fmt.Printf("List of bookings : %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will get the confirmation via email %v\n", firstName, lastName, userTickets, email)
	fmt.Println("Total slice:", bookings)
	fmt.Printf("Type of slice %T\n", bookings)
	fmt.Printf("Remaining tickets are %v\n", remainingTickets)
}
func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets are booked for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Ticket %v is sending to gmail %v\n", ticket, email)
	fmt.Println("#################")
	wait.Done()

}
