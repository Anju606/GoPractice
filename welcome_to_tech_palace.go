/*
There is an appliance store called "Tech Palace" nearby. The owner of the store recently installed a big display to use for marketing messages and to show a special greeting when customers scan their loyalty cards at the entrance. The display consists of lots of small LED lights and can show multiple lines of text.

The store owner needs your help with the code that is used to generate the text for the new display.
*/
package techpalace

import "strings"

const greeting = "Welcome to the Tech Palace"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	greetings := greeting + ", " + strings.ToUpper(customer)
	return greetings
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	greetings := strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
	return greetings
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	greetings := strings.Replace(oldMsg, "*", "", -1)
	greetings = strings.TrimSpace(greetings)
	return greetings
}
