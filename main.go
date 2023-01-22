//https://www.youtube.com/watch?v=yyUHQIec83I
//Golang Tutorial for Beginners | Full Go Course - TechWorld with Nana
//1:11:13 - Loops in Go

package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference" // ou só conferenceName := "Go Conference", não funciona pra const
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avalaible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//var bookings = [50]string{} //tem q botar o número máximo de itens e o tipo. No caso, 50 strings
	var bookings []string

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)//o "&" indica ponteiro

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)//o "&" indica ponteiro

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)//o "&" indica ponteiro

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)//o "&" indica ponteiro

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings { //blank identifier "_" ignore a variable you don't want to use
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all our bookings: %v\n", firstNames)


		//var noTicketsRemaing bool = remainingTickets == 0
		//noTicketsRemaing := remainingTickets == 0
		//if noTicketsRemaing { } 
		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}
}

/*slice is an abstraction of an array
more flexible, resizable
array = [], no value means it's a slice
array = append(array, value added)*/
//range iterates over elements for differents data structures
//For arrays and slices, range provides the index and value for each element