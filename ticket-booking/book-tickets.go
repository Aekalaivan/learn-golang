package main

import (
	"fmt"
)

func main() {
	const totalNumberOfSlots uint = 50
	var remainingTickets uint = totalNumberOfSlots

	fmt.Println("This is simple program for ticket booking")
	fmt.Printf("Total of %v seats available now\n", totalNumberOfSlots)

	for remainingTickets <= totalNumberOfSlots && remainingTickets != 0 {
		remainingTickets = booktTickets(remainingTickets)
	}
	fmt.Println("All tickets got booked!!!")
}

func booktTickets(remainingTickets uint) uint {

	fmt.Println("Please enter the number of tickets you would like to book")

	var numberOfTickets uint
	fmt.Scan(&numberOfTickets)

	if numberOfTickets <= 0 {
		fmt.Println("Invalid input, Please enter a valid input")
		return remainingTickets
	} else if numberOfTickets > remainingTickets {
		fmt.Printf("Sorry, Only %v seats remaining\n", remainingTickets)
		return remainingTickets
	} else {
		fmt.Println("Thanks for booking!")
		return remainingTickets - numberOfTickets
	}
}
