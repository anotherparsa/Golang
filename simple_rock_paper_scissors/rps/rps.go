package rps

import (
	"math/rand"
	"net/http"
)

type Round struct {
	Winner          string `json:"winner"`
	Computer_choice string `json:"computer_choice"`
	User_choice     string `json:"user_choice"`
}

func Play_round(w http.ResponseWriter, r *http.Request) Round {
	choices := []string{"rock", "paper", "scissors"}
	var user_choice string
	computer_choice := choices[rand.Intn(len(choices))]
	winner := Determine_winner(user_choice, computer_choice)
	round := Round{
		Winner:          winner,
		Computer_choice: computer_choice,
		User_choice:     user_choice,
	}
	return round
}

func Determine_winner(user_choice, computer_choice string) string {
	if user_choice == computer_choice {
		return "tie"
	}
	switch user_choice {
	case "rock":
		if computer_choice == "scissors" {
			return "user"
		}
		return "computer"
	case "paper":
		if computer_choice == "rock" {
			return "user"
		}
		return "computer"
	case "scissors":
		if computer_choice == "paper" {
			return "computer"
		}
		return "computer"
	default:
		return "err"
	}
}
