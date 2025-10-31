package main

import (
	"fmt"
	"strings"
)

func namevalidation(fn string, ln string) bool {
	if len(fn) < 2 || len(ln) < 2 {
		fmt.Println("First name and last name must be at least two characters long. Try again.")
		return false
	}
	return true
}

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Println("Welcome to the", conferenceName, "booking application!")
	fmt.Println("There are", remainingTickets,"out of", conferenceTickets, "tickets left.")
	fmt.Println("Get your tickets here to attend.")

	var bookings = []string{}

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("\nEnter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("\nEnter your last name: ")
		fmt.Scan(&lastName)
		if namevalidation(firstName, lastName) == false {
			continue
		}

		fmt.Println("\nEnter your email: ")
		fmt.Scan(&email)
		if ! strings.Contains(email, "@") {
			fmt.Println("Please enter a valid email address format.")
			continue
		}

		fmt.Println("\nEnter number of tickets: ")
		fmt.Scan(&userTickets)


		if userTickets > remainingTickets{
			fmt.Printf("There are %v tickets remaining. Please try a smaller number.\n\n", remainingTickets)
			continue
		}

		fmt.Printf("\nThank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
		fmt.Printf("You will receive a confirmation email at %v\n\n", email)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("%v out of %v tickets remaining.\n", remainingTickets, conferenceTickets)

		firstNames := []string{}
		for _, booking := range bookings {

			// Split each element of the array on a space.
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("\nHere are all the first names who booked: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("No more tickets available. Thank you for playing!")
			break
		}

	}
}
