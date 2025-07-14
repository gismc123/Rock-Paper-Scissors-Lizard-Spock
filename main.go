package main // Declare the main package

import (
	"bufio"     // for buffered input from the user
	"fmt"       // for formatted I/O like Println
	"math/rand" // For generating random numbers (bot's choice)
	"os"        // For OS-related functions, like reading input
	"strings"   // For string manipulation (trimming, lowercasing)
	"time"      // For seeding the random number generator
	"strconv"   // string-to-int conversion.
)

func main() {

	fmt.Printf("----------------------\n")
	
	fmt.Println("Starting a game of Rock Paper Scissors Lizard Spock..") // Print welcome message
	fmt.Println("first to 5 wins the game!")							  // Explain the winning condition

	fmt.Printf("----------------------\n \n")

	rand.Seed(time.Now().UnixNano()) // Seed random number generator with current time

	// Array holding the possible bot choices
	choiceOptions := [5]string{"rock","paper","scissors","lizard","spock"}

	// Initialize Score Board
	playerScore := 0 // Set inital score to zero
	botScore := 0    // Set inital score to zero
	drawScore := 0   // Set inital score to zero
	
	// Print initial scores
    fmt.Printf("Player score: %d \n", playerScore)
    fmt.Printf("Bot score: %d \n", botScore)
	fmt.Printf("Draw score: %d \n", drawScore)

	fmt.Printf("\n ----------------------\n \n")

	reader := bufio.NewReader(os.Stdin) // create a buffered reader to read user input

	for { 
		fmt.Printf("Choose your pick. Enter the number corresponding to your choice.\n Type Exit at anytime to quit the game. \n \n") // Request for input
			// for loop to provide the user with options
			for i, a := range choiceOptions { 
				fmt.Println (i,":", a)
			}
		answer, _ := reader.ReadString('\n')  // Read input until newline
		text := strings.TrimSpace(strings.ToLower(answer)) // Trim whitespace and convert to lowercase

		if text == "exit" { // If user types "exit", quit the game
				fmt.Println("Exiting the game...")
				break
		} else if text == "" { // If input is empty, warn user and continue loop
				fmt.Println("Empty input detected")
				continue
		} else if text != "0" && text != "1" && text != "2" && text != "3" && text != "4" { // Validate input
				fmt.Println("Please choose a number for either rock, paper, scissors, lizard or spock!")
				continue
		}

		// Generate a random index to pick bot's choice
		randomIndex := rand.Intn(len(choiceOptions))
		botPick := choiceOptions[randomIndex] // Bot's choice

		// Determine the winner and update scores accordingly
		index, err := strconv.Atoi(text)
			if err != nil || index < 0 || index >= len(choiceOptions) {
				fmt.Println("Invalid input. Please enter a number between 0 and 4.")
				continue
			}
			yourPick := choiceOptions[index]

		fmt.Printf("\n")
		// Show choices
		fmt.Println("You chose " + yourPick)
		fmt.Println("Bot chose " + botPick)

		fmt.Printf("----------------------\n")

		if botPick == "scissors" && yourPick == "paper"{
			fmt.Println("Scissors cuts paper.")
			fmt.Println("You lose!")
			botScore++
		} else if botPick == "scissors" && yourPick == "lizard"{
			fmt.Println("Scissors deacpitates lizard.")
			fmt.Println("You lose!")
			botScore++
		} else if botPick == "scissors" && yourPick == "rock"{
			fmt.Println("Rock crushes lizard.")
			fmt.Println("You win!")
			playerScore++
		} else if botPick == "scissors" && yourPick == "spock"{
			fmt.Println("Spock smashes scissors.")
			fmt.Println("You win!")
			playerScore++
		} else if botPick == "rock" && yourPick == "paper"{
			fmt.Println("Paper covers rock.")
			fmt.Println("You win!")
			playerScore++
		} else if botPick == "rock" && yourPick == "lizard"{
			fmt.Println("Rock crushes lizard.")
			fmt.Println("You lose!")
			botScore++
		} else if botPick == "rock" && yourPick == "spock"{
			fmt.Println("Spock vaporizes rock.")
			fmt.Println("You win!")
			playerScore++
		} else if botPick == "rock" && yourPick == "scissors"{
			fmt.Println("Rock crushes scissors.")
			fmt.Println("You lose!")
			botScore++
		} else if botPick == "paper" && yourPick == "lizard"{
			fmt.Println("Lizard eats paper.")
			fmt.Println("You win!")
			playerScore++
		} else if botPick == "paper" && yourPick == "spock"{
			fmt.Println("Paper disproves Spock.")
			fmt.Println("You lose!")
			botScore++
		} else if botPick == "paper" && yourPick == "scissors"{
			fmt.Println("Scissors cuts paper.")
			fmt.Println("You win!")
			playerScore++
		} else if botPick == "paper" && yourPick == "rock"{
			fmt.Println("Paper covers rock.")
			fmt.Println("You lose!")
			botScore++
		} else if botPick == "lizard" && yourPick == "spock"{
			fmt.Println("Lizard poisons Spock.")
			fmt.Println("You lose!")
			botScore++
		} else if botPick == "lizard" && yourPick == "scissors"{
			fmt.Println("Scissors decapitates lizard.")
			fmt.Println("You win!")
			playerScore++
		} else if botPick == "lizard" && yourPick == "rock"{
			fmt.Println("Rock crushes lizard.")
			fmt.Println("You win!")
			playerScore++
		} else if botPick == "lizard" && yourPick == "paper"{
			fmt.Println("Lizard eats paper.")
			fmt.Println("You lose!")
			botScore++
		} else if botPick == "spock" && yourPick == "scissors"{
			fmt.Println("Spock smashes scissors.")
			fmt.Println("You lose!")
			botScore++
		} else if botPick == "spock" && yourPick == "rock"{
			fmt.Println("Spock vaporizes rock.")
			fmt.Println("You lose!")
			botScore++
		} else if botPick == "spock" && yourPick == "paper"{
			fmt.Println("Paper Disproves Spock.")
			fmt.Println("You win!")
			playerScore++
		} else if botPick == "spock" && yourPick == "lizard"{
			fmt.Println("Lizard poisons Spock.")
			fmt.Println("You win!")
			playerScore++
		} else {
			// If none of the above, it's a draw
			fmt.Println("Draw!")
			drawScore++
        }

		fmt.Printf("----------------------\n")

		// Print updated scores
		fmt.Printf("Player score: %d \n", playerScore)
		fmt.Printf("Bot score: %d \n", botScore)
		fmt.Printf("Draw score: %d \n", drawScore)

		fmt.Printf("----------------------\n")

		// Check if either player reached 5 points to end the game
		if playerScore >= 5 {
			fmt.Println("You won the game! Congratulations!")
				break
		} else if botScore >= 5 {
			fmt.Println("You lose the game! Better luck next time!")
				break
		}
		
	}

}