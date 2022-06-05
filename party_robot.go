/* __________PROBLEM STATEMENT_______________
Once there was an eccentric programmer living in a strange house with barred windows. One day he accepted a job from an online job board to build a party robot. The robot is supposed to greet people and help them to their seats. The first edition was very technical and showed the programmer's lack of human interaction. Some of which also made it into the next edition.

1. Implement the Welcome function to return a welcome message using the given name:

2. Implement the HappyBirthday function to return a birthday message using the given name and age of the person. Unfortunately the programmer is a bit of a show-off, so the robot has to demonstrate its knowledge of every guest's birthday.

3. Implement the AssignTable function to give directions. It should accept 5 parameters.
 - The name of the new guest to greet (string)
 - The table number (int)
 - The name of the seatmate (string)
 - The direction where to find the table (string)
 - The distance to the table (float64)
The robot provides the table number in a 3 digits format. If the number is less than 3 digits it gets extra leading zeroes to become 3 digits (e.g. 3 becomes 003). The robot also mentions the distance of the table. Fortunately only with a precision that's limited to 1 digit.
    For Example: AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298)
                   // Output:
                   // Welcome to my party, Christiane!
                   // You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
                   // You will be sitting next to Frank.

*/

package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	part1 := Welcome(name)
	part2 := fmt.Sprintf("\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.", table, direction, distance)
	part3 := fmt.Sprintf("\nYou will be sitting next to %s.", neighbor)
	return part1 + part2 + part3
}
