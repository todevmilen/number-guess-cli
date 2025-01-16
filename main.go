package main

import (
	"math/rand/v2"
)

import (
	"fmt"
	"log"
)

func main() {
	difficultyLevels := []string{"Easy", "Medium", "Hard"}
	chances := []int{10, 5, 3}

	fmt.Println(`
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.


Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
		`)

	var difficulty int
	var guess int

	fmt.Print("Enter your choice: ")
	d, err := fmt.Scanln(&difficulty)
	if err != nil || d != 1 {
		log.Fatal("Please enter valid choice 1, 2 or 3")
	}

	if difficulty != 1 && difficulty != 2 && difficulty != 3 {
		log.Fatal("Please enter valid choice 1, 2 or 3")
	}

	fmt.Printf(`
Great! You have selected the %s difficulty level.
Let's start the game!
`, difficultyLevels[difficulty-1])

	randomNumber := rand.IntN(100-1) + 1

	currChance := 0

	for currChance < chances[difficulty-1] {
		fmt.Print("Enter your guess: ")
		c, err := fmt.Scanln(&guess)
		if err != nil || c != 1 {
			fmt.Println("Please enter a valid integer")
			currChance += 1
			continue
		}
		if guess < 0 || guess > 100 {
			fmt.Println("You have to guess between 1 and 100")
			currChance += 1
			continue
		}
		if guess > randomNumber {
			fmt.Printf("Incorrect! The number is less than %v\n", guess)
			currChance += 1
			continue
		}

		if guess < randomNumber {
			fmt.Printf("Incorrect! The number is greater than %v\n", guess)
			currChance += 1
			continue
		}

		if guess == randomNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %v attempts\n", currChance)
			break
		}
	}

}
