/*In this exercise you'll be writing code to analyze the production in a car factory.

1. The cars are produced on an assembly line. The assembly line has a certain speed, that can be changed. The faster the assembly line speed is, the more cars are produced. However, changing the speed of the assembly line also changes the number of cars that are produced successfully, that is cars without any errors in their production. Implement a function that takes in the number of cars produced per hour and the success rate and calculates the number of successful cars made per hour. The success rate is given as a percentage, from 0 to 100.

2. Implement a function that takes in the number of cars produced per hour and the success rate and calculates how many cars are successfully produced each minute.

3. Each car normally costs $10,000 to produce individually, regardless of whether it is successful or not. But with a bit of planning, 10 cars can be produced together for $95,000. For example, 37 cars can be produced in the following way: 37 = 3 x groups of ten + 7 individual cars. So the cost for 37 cars is: 3*95,000+7*10,000=355,000. Implement the function CalculateCost that calculates the cost of producing a number of cars, regardless of whether they are successful */

package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var res float64 = float64(productionRate) * successRate / 100
	return res
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var res = CalculateWorkingCarsPerHour(productionRate, successRate)
	result := int(res / 60)
	return result
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	tens := int(carsCount / 10)
	ones := carsCount % 10
	return uint(tens*95000 + ones*10000)
}
