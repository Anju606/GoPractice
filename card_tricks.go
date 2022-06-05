/* ____________________PROBLEM STATEMENT(SLICE)____________________
As a magician-to-be, Elyse needs to practice some basics. She has a stack of cards that she wants to manipulate. To make things a bit easier she only uses the cards 1 to 10.

1. When practicing with her cards, Elyse likes to start with her favorite three cards of the deck: 2, 6 and 9. Write a function FavoriteCards that returns a slice with those cards in that order.

2. Return the card at position index from the given stack. If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to return -1:

3. Exchange the card at position index with the new card provided and return the adjusted stack. Note that this will modify the input slice which is the expected behavior. If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to append the new card to the end of the stack.

4. Add the card(s) specified in the value parameter at the top of the stack. If no argument is given for the value parameter, then the result equals the original slice.

5. Remove the card at position index from the stack and return the stack. Note that this may modify the input slice which is ok. If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to leave the stack unchanged:

*/
package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	var sl = []int{2, 6, 9}
	return sl
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index int, value int) []int {

	if index == len(slice) || index == -1 {
		slice = append(slice, value)
		return slice
	}
	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
	slice = append(value, slice...)
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	ind := index + 1
	if index > len(slice) || index < 0 {
		return slice
	}
	slice = append(slice[:index], slice[ind:]...)
	return slice
}
