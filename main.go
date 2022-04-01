package main

import (
	"fmt"
	"sync"
	"time"
	// "strconv"
)

// var aArray = [3]string{"dsfs", "dsfds", "sdfsdf"}

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingtickets uint = 50

type UserData struct {
	firstName   string
	lastName    string
	emailId     string
	noOfTickets uint
}

// var aBookings = [50]string{"RAM", "KITU"}
// var aBookings []string // slice
var aTempMaps = make([]map[string]string, 0)

var aBookings []UserData

// wait group used to wait(hold) main thread upto complte sub threads
var oWait = sync.WaitGroup{}

func main() {
	var sDoForMultiRequests string
	fmt.Println("Do you want to enter multiple requests: Y/N")
	fmt.Scan(&sDoForMultiRequests)
	if sDoForMultiRequests == "Y" || sDoForMultiRequests == "y" {
		multiRequestProcessConcurrencyGoRouting()
	} else {
		singleProcessSynchoronizeGoRuouting()
	}
}

func multiRequestProcessConcurrencyGoRouting() {
	greetUsers()
	fmt.Printf("types of: Conference name is %T , ConferenceTickets is %T, remaininfTickets is %T \n", conferenceName, conferenceTickets, remainingtickets)

	for remainingtickets > 0 && len(aBookings) < 50 {
		sFirstName, sLastName, sEmailId, iNoOfTickets := getUserInputs()

		if !validateUserInputs(sFirstName, sLastName, sEmailId, iNoOfTickets) {
			continue
		}
		bookTickets(sFirstName, sLastName, sEmailId, iNoOfTickets)
		go sendTickets(sFirstName, sLastName, sEmailId, iNoOfTickets, false)

		printFirstNames()

		// bHasTicketsSoldOut := remainingtickets == 0
		if remainingtickets == 0 {
			fmt.Printf("Our Conference is booked out. come back next year.")
			break
		}
		// }
	}
}

func singleProcessSynchoronizeGoRuouting() {
	greetUsers()
	fmt.Printf("types of: Conference name is %T , ConferenceTickets is %T, remaininfTickets is %T \n", conferenceName, conferenceTickets, remainingtickets)

	sFirstName, sLastName, sEmailId, iNoOfTickets := getUserInputs()

	if !validateUserInputs(sFirstName, sLastName, sEmailId, iNoOfTickets) {
		return
	}

	bookTickets(sFirstName, sLastName, sEmailId, iNoOfTickets)

	oWait.Add(1)                                                        // add waiting list for sub thread
	go sendTickets(sFirstName, sLastName, sEmailId, iNoOfTickets, true) // go is used to create sub threads

	printFirstNames()

	if remainingtickets == 0 {
		fmt.Printf("Our Conference is booked out. come back next year.")
		// break
	}

	oWait.Wait() // wait here upto complete waiting list
}

func greetUsers() {
	fmt.Println("Hello there, Welcome to", conferenceName, "Booking app site")
	fmt.Printf("we have totoal of %v tickets and %v are still available", conferenceTickets, remainingtickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() {
	fmt.Printf("The first names of bookings are: %v \n", getFirstNames(aBookings))
}

func getFirstNames(aBookings []UserData) []string {
	aFirstNames := []string{}
	for _, booking := range aBookings {
		aFirstNames = append(aFirstNames, booking.firstName)
	}
	return aFirstNames
}

func getUserInputs() (string, string, string, uint) {
	var sFirstName string
	var sLastName string
	var sEmailId string
	var iNoOfTickets uint
	// var iUserTickets int
	fmt.Println("Enter your first name:")
	fmt.Scan(&sFirstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&sLastName)

	fmt.Println("Enter your email id:")
	fmt.Scan(&sEmailId)

	fmt.Println("Enter no.of tickets:")
	fmt.Scan(&iNoOfTickets)

	return sFirstName, sLastName, sEmailId, iNoOfTickets
}

func bookTickets(sFirstName string, sLastName string, sEmailId string, iNoOfTickets uint) {
	remainingtickets = remainingtickets - uint(iNoOfTickets)
	// aBookings[0] = sFirstName + " " + sLastName
	// aBookings = append(aBookings, sFirstName+" "+sLastName)

	// var myMap map[string]string
	// var oUserMap = make(map[string]string)
	// oUserMap["firstName"] = sFirstName
	// oUserMap["lastName"] = sLastName
	// oUserMap["emailId"] = sEmailId
	// oUserMap["noOfTickets"] = strconv.FormatUint(uint64(iNoOfTickets), 10)
	// aTempMaps = append(aTempMaps, oUserMap)

	// Use struct
	var oUserData = UserData{
		firstName:   sFirstName,
		lastName:    sLastName,
		emailId:     sEmailId,
		noOfTickets: iNoOfTickets,
	}
	aBookings = append(aBookings, oUserData)

	fmt.Printf("The whole slice %v \n", aBookings)
	fmt.Printf("The first slice %v \n", aBookings[0])

	// sUserName = "Ram"
	// iUserTickets = 2
	fmt.Printf("Thank you %v %v for booking %v tickets tickets. You will receive a confirmation mail at %v \n", sFirstName, sLastName, iNoOfTickets, sEmailId)
	fmt.Printf("%v remaining tickets of %v \n", remainingtickets, conferenceTickets)
}

func sendTickets(sFirstName string, sLastName string, sEmailId string, iNoOfTickets uint, bWaiting bool) {
	// wait for 10 seconds to process
	time.Sleep(10 * time.Second)
	var sTicketsInfo = fmt.Sprintf("%v tickets are booked by name %v %v", iNoOfTickets, sFirstName, sLastName)
	fmt.Println("#######################")
	fmt.Printf("Sending tickets:\n %v \n to email addtress %v \n", sTicketsInfo, sEmailId)
	fmt.Println("#######################")
	if bWaiting {
		oWait.Done() // Sub Thread process complete, set waiting queue entity as done.
	}
}
