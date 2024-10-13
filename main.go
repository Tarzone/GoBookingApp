package main

import "fmt"

func main() {
	// var conferenceName = "Go conference"
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 40

	// fmt.Println("Welcome to", conferenceName, "Hello World")
	// fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	// fmt.Println("Get yout tickets here to attend.")

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend.\n")

	var userName string
	var userTickets int
	//  ask user for their name
	fmt.Scan()

	userName = "Isaac"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
