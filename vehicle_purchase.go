/* ____________________PROBLEM STATEMENT(COMPARISON)____________________
You have three tasks, one to determine if you need a license, one to help you choose between two vehicles and one to estimate the acceptable price for a used vehicle.

1. Some vehicle kinds require a driver's license to operate them. Assume only the kinds "car" and "truck" require a license, everything else can be operated without a license. Implement the NeedsLicense(kind) function that takes the kind of vehicle and returns a boolean indicating whether you need a license for that kind of vehicle.

2. You evaluate your options of available vehicles. You manage to narrow it down to two options but you need help making the final decision. For that, implement the function ChooseVehicle(option1, option2) that takes two vehicles as arguments and returns a decision that includes the option that comes first in lexicographical order.

3. Now that you made a decision, you want to make sure you get a fair price at the dealership. Since you are interested in buying a used vehicle, the price depends on how old the vehicle is. For a rough estimate, assume if the vehicle is less than 3 years old, it costs 80% of the original price it had when it was brand new. If it is at least 10 years old, it costs 50%. If the vehicle is at least 3 years old but less than 10 years, it costs 70% of the original price.
*/
package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	}
	return option2 + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		rate := 0.8
		return rate * originalPrice
	}
	if age >= 10 {
		rate := 0.5
		return rate * originalPrice
	}
	rate := 0.7
	return rate * originalPrice
}
