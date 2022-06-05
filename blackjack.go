/* ____________________PROBLEM STATEMENT(SWITCH)____________________
In this exercise we will simulate the first turn of a Blackjack game.
You will receive two cards and will be able to see the face up card of the dealer. All cards are represented using a string such as "ace", "king", "three", "two", etc. The values of each card are:
card	value	card	value
ace	     11	    eight	  8
two    	  2	    nine	  9
three	  3	     ten	 10
four	  4	    jack	 10
five	  5	    queen	 10
six	      6	     king	 10
seven	  7	    other	  0

Note: Commonly, aces can take the value of 1 or 11 but for simplicity we will assume that they can only take the value of 11.
Depending on your two cards and the card of the dealer, there is a strategy for the first turn of the game, in which you have the following options:

 - Stand (S)
 - Hit (H)
 - Split (P)
 - Automatically win (W)

Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

 - If you have a pair of aces you must always split them.
 - If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
 - If your cards sum up to a value within the range [17, 20] you should always stand.
 - If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
 - If your cards sum up to 11 or lower you should always hit.

1. Implement a function to calculate the numerical value of a card:

2. Write a function that implements the decision logic as described above:

*/
package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int = 0
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten":
		value = 10
	case "jack":
		value = 10
	case "queen":
		value = 10
	case "king":
		value = 10
	default:
		value = 0
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var value1 int = ParseCard(card1)
	var value2 int = ParseCard(card2)
	var sum int = value1 + value2
	var value3 int = ParseCard(dealerCard)
	switch {
	case sum == 22:
		return "P"
	case sum == 21 && value3 < 10 && value3 != 0:
		return "W"
	case sum == 21 && (value3 >= 10 || dealerCard != "joker"):
		return "S"
	case sum <= 20 && sum >= 17:
		return "S"
	case (sum <= 16 && sum >= 12) && value3 >= 7:
		return "H"
	case (sum <= 16 && sum >= 12) && value3 < 7:
		return "S"
	case sum <= 11:
		return "H"
	}
	return "W"
}
