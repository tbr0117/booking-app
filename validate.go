package main

import (
	"fmt"
	"booking-app/helper"
)


func validateUserInputs(sFirstName string, sLastName string, sEmailId string, iNoOfTickets uint) bool {
	isValidName, isValidEmail, isValidTicketNo := helper.ValidateValues(sFirstName, sLastName, sEmailId, iNoOfTickets, remainingtickets)
	if !isValidName || !isValidEmail {
		fmt.Printf("Your Input data is not valid \n")
	} else if !isValidTicketNo {
		fmt.Printf("numbers tickets must be greterthan 0 \n")
	} else if iNoOfTickets > remainingtickets {
		fmt.Printf("We have only %v tickets available, so you can't book %v tickets /\n", remainingtickets, iNoOfTickets)
	}
	return isValidName && isValidEmail && isValidTicketNo && iNoOfTickets <= remainingtickets
}