package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	target := random.Intn(100) + 1
	var userGuess int

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a number between 1 and 100. Can you guess it?")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&userGuess)

		if userGuess == target {
			fmt.Println("Congratulations! You've guessed the correct number!")
			break
		} else if userGuess < target {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}
}
