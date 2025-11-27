package helper

import "strings"

func ValidAte(frsN string, lstN string, emAil string, userTc int, remainingTickets int) (bool, bool, bool) {
	validName := len(frsN) >= 2 && len(lstN) >= 2
	validEmail := strings.Contains(emAil, "@")
	validTicketNumber := userTc > 0 && userTc <= remainingTickets
	return validName, validEmail, validTicketNumber
}
