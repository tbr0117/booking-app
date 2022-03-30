package helper

import (
	"strings"
)

func ValidateValues(sFirstName string, sLastName string, sEmailId string, iNoOfTickets uint, remainingtickets uint) (bool, bool, bool) {
	var isValidName bool = len(sFirstName) >= 2 && len(sLastName) >= 2
	isValidEmail := strings.Contains(sEmailId, "@")
	isValidTicketNo := iNoOfTickets > 0 && iNoOfTickets < remainingtickets
	return isValidName, isValidEmail, isValidTicketNo
}