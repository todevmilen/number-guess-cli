# Number Guessing Game

Welcome to the **Number Guessing Game**! This is a simple and fun game where the program generates a random number between 1 and 100, and your goal is to guess the number within a limited number of chances based on the selected difficulty level.

## How It Works

1. **Start the Game**:
   - The program will welcome you and prompt you to choose a difficulty level.
2. **Select Difficulty**:
   - `1. Easy` (10 chances)
   - `2. Medium` (5 chances)
   - `3. Hard` (3 chances)
3. **Guess the Number**:
   - Enter your guesses one at a time.
   - The program will provide feedback if the guessed number is higher or lower than the target number.
4. **Win or Lose**:
   - If you guess the number correctly within the allowed chances, you win.
   - If you run out of chances, the game ends and the number is revealed.

## Features

- Difficulty levels with different numbers of chances.
- Input validation to ensure valid guesses.
- Feedback on whether the guess is too high or too low.
- Randomized target number to keep the game challenging.

## Requirements

- Go programming language installed (minimum version 1.19).

## How to Run

1. Clone the repository or copy the code.
2. Save the code in a file named `main.go`.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the program using the following command:
   ```bash
   go run main.go
   ```
5. Follow the on-screen instructions to play the game.

## Code Overview

### Packages Used

- **`math/rand/v2`**: Generates a random number for the game.
- **`fmt`**: Handles input and output operations.
- **`log`**: Logs errors and handles invalid inputs.

### Main Logic

- The program initializes difficulty levels and chances.
- It prompts the user to select a difficulty level.
- A random number between 1 and 100 is generated.
- The user guesses the number, with feedback provided after each guess.
- The game ends either when the user guesses correctly or when they run out of chances.

## Example Output

```
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice: 2

Great! You have selected the Medium difficulty level.
Let's start the game!

Enter your guess: 50
Incorrect! The number is less than 50
Enter your guess: 25
Incorrect! The number is greater than 25
Enter your guess: 30
Congratulations! You guessed the correct number in 3 attempts
```

## This project is part of [RoadMapSH Projects](https://roadmap.sh/projects/number-guessing-game)
