/* ____________________PROBLEM STATEMENT(RANDOMNESS)____________________
Elaine is working on a new children's game that features animals and magic wands. It is time to code functions for rolling a die, generating random wand energy and shuffling a slice.

1. Implement a SeedWithTime function that seeds the math.rand package with the current computer time.

2. Implement a RollADie function. This will be the traditional twenty-sided die with numbers 1 to 20.

3. Implement a GenerateWandEnergy function. The wand energy should be a random floating point number between 0.0 and 12.0.

4. The game features eight different animals:
 - ant
 - beaver
 - cat
 - dog
 - elephant
 - fox
 - giraffe
 - hedgehog
Write a function ShuffleAnimals that returns a slice with the eight animals in random order.

*/
package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	v := rand.Intn(20) + 1
	return v
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	v := rand.Float64() * 12.0
	return v
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	v := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(v), func(i, j int) {
		v[i], v[j] = v[j], v[i]
	})

	return v
}
