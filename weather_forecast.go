/* Being hired by a big weather forecast company, your job is to maintain their code base. Scrolling through the code, you find it hard to follow what the code is actually doing. Good thing you know just the right thing to do!

1. Write a package comment for package weather that describes its contents. The package comment should introduce the package and provide information relevant to the package as a whole.

2. Write a comment for the package variable CurrentCondition. This should tell any user of the package what information the variable stores, and what they can do with it.

3. Just like the previous step, write a comment for CurrentLocation.

4. Write a function comment for Forecast(). Your comments must describe what the function does, but not how it does it.*/

// Package weather ...
package weather

// CurrentCondition variable contains the current condition of a location.
var CurrentCondition string

// CurrentLocation variable contains the current location of interest.
var CurrentLocation string

// Forecast function prints the current location and its current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
