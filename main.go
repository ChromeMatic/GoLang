package main
import "fmt"

func main() {

	var conferenceName string = "GO CONFERENCE"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Println("================================================")
	fmt.Printf("   WELCOME TO %v BOOKING APPLICATION \n",conferenceName)
	fmt.Printf("   WE HAVE TICKETS %v FOR THE CONFERENCE \n",conferenceTickets)
	fmt.Printf("   AND WE HAVE %v TICKETS REMAINING.\n",remainingTickets)
	fmt.Println("================================================")

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

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("\n\tUser %v %v booked %v tickets.\n\tYou will receive a conformation email at %v", firstName, lastName, userTickets, email)
	
	fmt.Println("")
	fmt.Println("\n\t================================================")
	fmt.Printf("\t= Remaining tickets for %v is %v tickets \n",conferenceName,remainingTickets)
	fmt.Println("\t================================================")
}
