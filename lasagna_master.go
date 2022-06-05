/* ____________________PROBLEM STATEMENT(FUNCTIONS)____________________
In this exercise you are going to write some more code related to preparing and cooking your brilliant lasagna from your favorite cookbook. You have four tasks. The first one is related to the cooking itself, the other three are about the perfect preparation.

1. For the next lasagna that you will prepare, you want to make sure you have enough time reserved so you can enjoy the cooking. You already planned which layers your lasagna will have. Now you want to estimate how long the preparation will take based on that. Implement a function PreparationTime that accepts a slice of layers as a []string and the average preparation time per layer in minutes as an int. The function should return the estimate for the total preparation time based on the number of layers as an int. Go has no default values for functions. If the average preparation time is passed as 0 (the default initial value for an int), then the default value of 2 should be used.

2. Besides reserving the time, you also want to make sure you have enough sauce and noodles to cook the lasagna of your dreams. For each noodle layer in your lasagna, you will need 50 grams of noodles. For each sauce layer in your lasagna, you will need 0.2 liters of sauce. Define the function Quantities that takes a slice of layers as parameter as a []string. The function will then determine the quantity of noodles and sauce needed to make your meal. The result should be returned as two values of noodles as an int and sauce as a float64.

3. A while ago you visited a friend and ate lasagna there. It was amazing and had something special to it. The friend sent you the list of ingredients and told you the last item on the list is the "secret ingredient" that made the meal so special. Now you want to add that secret ingredient to your recipe as well. Write a function AddSecretIngredient that accepts two slices of ingredients of type []string as parameters. The first parameter is the list your friend sent you, the second is the ingredient list of your own recipe. The last element in your ingredient list is always "?". The function should replace it with the last item from your friends list. Note: AddSecretIngredient does not return anything - you should modify the list of your ingredients directly. The list with your friend's ingredients should not be modified. Also, since slice is passed into a function as pointers, changes to the two []string arguments passed into AddSecretIngredient will be modified directly.

4. The amounts listed in your cookbook only yield enough lasagna for two portions. Since you want to cook for more people next time, you want to calculate the amounts for different numbers of portions. Implement a function ScaleRecipe that takes two parameters.
 - A slice of float64 amounts needed for 2 portions.
 - The number of portions you want to cook.
The function should return a slice of float64 of the amounts needed for the desired number of portions. You want to keep the original recipe though. This means the amounts argument should not be modified in this function.

*/
package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}

	return time * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	i := 0
	noodles := 0
	sauce := 0.0
	for i = 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			sauce += 1
		}

		if layers[i] == "noodles" {
			noodles += 1
		}
	}

	return noodles * 50, sauce * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) []string {
	myList[len(myList)-1] = friendList[len(friendList)-1]
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	amount := make([]float64, len(amounts))
	i := 0
	for i = 0; i < len(amounts); i++ {
		amount[i] = amounts[i] * float64(portions) / 2
	}

	return amount
}
