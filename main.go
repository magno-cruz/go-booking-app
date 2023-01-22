//https://www.youtube.com/watch?v=yyUHQIec83I
//Golang Tutorial for Beginners | Full Go Course - TechWorld with Nana
//1:31:20 - Loops in Go

package main

import (
	"fmt"
	"go-booking-app/helper"
	"time"
	"sync"
)

var conferenceName = "Go Conference" // ou só conferenceName := "Go Conference", não funciona pra const
const conferenceTickets = 50
var remainingTickets uint = 50
//var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)
////package level variables = global variables. Have to use var
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}
//type creates custom types
//struct is like a map, but accept different types

var wg = sync.WaitGroup {}
//waits for the launched go routine to finish
//package sync provides basic synchronization functionality

func main() {
	

	greetUsers()

	//var bookings = [50]string{} //tem q botar o número máximo de itens e o tipo. No caso, 50 strings

	for /*remainingTickets > 0 && len(bookings) < 50*/ {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName,email)
			wg.Add(1)//number of goroutines the program has to wait for
			go sendTicket(userTickets, firstName, lastName,email)// goroutine, run this code on another thread and the rest of the code keeps going

			firstNames := getsFirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			//var noTicketsRemaing bool = remainingTickets == 0
			//noTicketsRemaing := remainingTickets == 0
			//if noTicketsRemaing { } 
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does't contain \"@\" sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid.")
			}
		}
	}
	wg.Wait()//blocks until the WaitGroup is (0)
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avalaible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getsFirstNames() []string{ //has to define the returning value's type
	firstNames := []string{}
	for _, booking := range bookings { //blank identifier "_" ignore a variable you don't want to use
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	//if I was using a map I would needed to convert the int

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sending ticket: \n%v \nto email address %v", ticket, email)
	fmt.Println("#######################")
	wg.Done()//decrements the waitgroup counter by 1, is called by the goroutine to indicate it is finished
}

/*slice is an abstraction of an array
more flexible, resizable
array = [], no value means it's a slice
array = append(array, value added)*/
//range iterates over elements for differents data structures
//For arrays and slices, range provides the index and value for each element
/*
city := "London"
switch city {
	case "London":
		//Code for London
	case "Singapore":
		//Code for Singapore
	case "New York", "Berlin":
		//Code for New York and Berlin is the same
	default:
		fmt.Println("No valid city selected")
}*/