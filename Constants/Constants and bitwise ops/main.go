package main

import "fmt"

const (
	emailAccess   = 1 << iota // 0000 0001 << 0 = 0000 0001
	office365                 // 0000 0001 << 1 = 0000 0010
	accessVPN                 // 0000 0001 << 2 = 0000 0100
	laptop                    // 0000 0001 << 3 = 0000 1000
	adminRights               // 0000 0001 << 4 = 0001 0000
	homeOffice                // 0000 0001 << 5 = 0010 0000
	payrollAccess             // 0000 0001 << 6 = 0100 0000
	traveler                  // 0000 0001 << 7 = 1000 0000
)

func main() {
	// Setting persmissions using OR
	var israRoles = emailAccess | accessVPN | laptop |
		adminRights | homeOffice | traveler // 1011 1101

	// Printing permissions
	fmt.Printf("Isra's permissions: %b\n\n", israRoles) // 1011 1001

	// Testing permissions using AND
	fmt.Printf("Does Isra have an email address? %v\n",
		yesNo(israRoles&emailAccess == emailAccess)) // Yes
	fmt.Printf("Does Isra have an O365 licence? %v\n",
		yesNo(israRoles&office365 == office365)) // No
	fmt.Printf("Does Isra have VPN access? %v\n",
		yesNo(israRoles&accessVPN == accessVPN)) // Yes
	fmt.Printf("Does Isra have a compnay laptop? %v\n",
		yesNo(israRoles&laptop == laptop)) // Yes
	fmt.Printf("Does Isra have a admin credentials? %v\n",
		yesNo(israRoles&adminRights == adminRights)) // Yes
	fmt.Printf("Can Isra work from home? %v\n",
		yesNo(israRoles&homeOffice == homeOffice)) // Yes
	fmt.Printf("Does Isra have payroll access? %v\n",
		yesNo(israRoles&payrollAccess == payrollAccess)) // No
	fmt.Printf("Does Isra travel for work? %v\n",
		yesNo(israRoles&traveler == traveler)) // Yes
}

func yesNo(b bool) string {
	// Returns string "Yes" or "No" based on the passed bool value
	if b {
		return "Yes"
	}
	return "No"
}
