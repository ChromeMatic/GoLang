package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName string = "GO CONFERENCE"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Println()
	fmt.Println("\t================================================")
	fmt.Printf("\t=   WELCOME TO %v BOOKING APPLICATION \n", conferenceName)
	fmt.Printf("\t=   WE HAVE TICKETS %v FOR THE CONFERENCE \n", conferenceTickets)
	fmt.Printf("\t=   AND WE HAVE %v TICKETS REMAINING.\n", remainingTickets)
	fmt.Println("\t================================================")

	//Using a slices
	var bookings []string

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Print("\tPlease enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("\n\tPlease enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("\n\tPlease enter your email address: ")
		fmt.Scan(&email)
		fmt.Print("\n\tPlease enter ticket amount you wish to buy: ")
		fmt.Scan(&userTickets)

		if userTickets > 50 {
			fmt.Printf("\n\t ===> Ticket amount %v exceed available tickets offered %v", userTickets, remainingTickets)
			fmt.Println()
			continue
		} else {
			remainingTickets = remainingTickets - userTickets
		}

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("\n\tUser %v booked %v tickets.\n\tYou will receive a conformation email at %v", bookings[0], userTickets, email)

		fmt.Println("")
		fmt.Println("\n\t===================================================")
		fmt.Printf("\t= Remaining tickets for %v is %v tickets \n", conferenceName, remainingTickets)
		fmt.Println("\t===================================================")
		fmt.Println("\t= All the books")

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("\t= %v \n", firstNames)
		fmt.Println()

		if remainingTickets == 0 {
			// end loop
			fmt.Println("\t==> Tickets are sold out! <===")
			break
		}
	}
}
