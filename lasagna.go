/* In this exercise you're going to write some code to help you cook a brilliant lasagna from your favorite cooking book.

You have four tasks, all related to the time spent cooking the lasagna.

1. Define the OvenTime constant with how many minutes the lasagna should be in the oven. According to the cooking book, the expected oven time in minutes is 40

2. Define the RemainingOvenTime() function that takes the actual minutes the lasagna has been in the oven as a parameter and returns how many minutes the lasagna still has to remain in the oven, based on the expected oven time in minutes from the previous task.

3. Define the PreparationTime function that takes the number of layers you added to the lasagna as a parameter and returns how many minutes you spent preparing the lasagna, assuming each layer takes you 2 minutes to prepare.

4. Define the ElapsedTime function that takes two parameters: the first parameter is the number of layers you added to the lasagna, and the second parameter is the number of minutes the lasagna has been in the oven. The function should return how many minutes in total you've worked on cooking the lasagna, which is the sum of the preparation time in minutes, and the time in minutes the lasagna has spent in the oven at the moment.*/

package lasagna

const OvenTime = 40 // TODO: define the 'OvenTime' constant

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return 40 - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
