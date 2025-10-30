package main


import "fmt"


func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Println("Welcome to the", conferenceName, "booking application!")
	fmt.Println("There are", remainingTickets,"out of", conferenceTickets, "tickets left.")
	fmt.Println("Get your tickets here to attend.")


	var firstName string
	var lastName string
	var email string
	var userTickets int
	// Ask user for their name

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)


	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation email at %v\n", email)

	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)

}
