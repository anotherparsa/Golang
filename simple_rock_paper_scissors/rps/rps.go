package rps

import (
	"math/rand"
	"net/http"
)

func play_round(w http.ResponseWriter, r *http.Request) {
	choices := []string{"rock", "paper", "scissors"}
	var userChoice string
	computerChoice := choices[rand.Intn(len(choices))]
	result := determineWinner(userChoice, computerChoice)
}

func determineWinner(user_choice, computer_choice string) string {
	if user_choice == computer_choice {
		return "It's a tie!"
	}

	switch user_choice {
	case "rock":
		if computer_choice == "scissors" {
			return "You win! Rock crushes scissors."
		}
		return "You lose! Paper covers rock."
	case "paper":
		if computer_choice == "rock" {
			return "You win! Paper covers rock."
		}
		return "You lose! Scissors cut paper."
	case "scissors":
		if computer_choice == "paper" {
			return "You win! Scissors cut paper."
		}
		return "You lose! Rock crushes scissors."
	default:
		return "Unexpected error."
	}

}
