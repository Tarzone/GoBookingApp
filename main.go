package main

import "fmt"

func main() {

	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 40

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend.\n")

	// var bookings = [50]string{"Rachel", "Wale", "Christine", "Maria"}
	// var bookings []string
	// var bookings = []string{}

	bookings := []string{}
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// var userName string
	// fmt.Scan(&userName)

	// Examplle outputs
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)
	// userName = "Isaac"
	// userTickets = 2

	//  ask user for their info
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	// fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
	// bookings[49] = firstName + " " + lastName

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The whole array: %v\n", bookings[0])
	// fmt.Printf("Slice type: %v\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings: %v\n", bookings)

}
