package main

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) {
	fmt.Print("Welcome to the Tech Palace, " + strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) {
	ash := strings.Repeat("*", numStarsPerLine)
	fmt.Print(ash + "\n" + welcomeMsg + "\n" + ash)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) {
	cleanmsg := strings.ReplaceAll(oldMsg, "*", "")
	fmt.Print(strings.TrimSpace(cleanmsg))

}
func main() {
	WelcomeMessage("Sumathy")
	CleanupMessage("***********Sumathy")
	AddBorder("Hi Sumathy!", 5)

}
