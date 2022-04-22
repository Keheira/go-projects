package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Action int64

// enum
const (
	Rock Action = iota
	Paper
	Scissors
)

// get enum readable
func (a Action) String() string {
	switch a {
	case Rock:
		return "rock"
	case Paper:
		return "paper"
	case Scissors:
		return "scissors"
	}
	return "hesitation"
}

func printAction(a Action) string {
	return a.String()
}

func checkUserWin(user Action, comp Action) bool {
	if (user == Rock && comp == Scissors) ||
		(user == Paper && comp == Rock) ||
		(user == Scissors && comp == Paper) {
		return true
	} else {
		return false
	}
}
func checkCompWin(user Action, comp Action) bool {
	if (comp == Rock && user == Scissors) ||
		(comp == Paper && user == Rock) ||
		(comp == Scissors && user == Paper) {
		return true
	} else {
		return false
	}
}

func getWinner(user Action, comp Action) {
	if user.String() == "hesitation" {
		fmt.Println("Sad way to lose friend.")
	} else if user == comp {
		fmt.Println("Tie. Break it up! Break it up!")
	} else if checkUserWin(user, comp) {
		fmt.Println("You won!")
	} else {
		fmt.Println("Sorry. That dang computer is too good.")
	}
}

func main() {
	var choice int
	rand.Seed(time.Now().UnixNano()) // not 100% why we have to seed
	var compChoice = rand.Intn(3)

	fmt.Println("Let's play a game")
	fmt.Print(
		"Rock, Paper, Scissors\n",
		"Input 0 for Rock\n",
		"Input 1 for Paper\n",
		"Input 2 for Scissors\n",
		"your choice: ",
	)
	fmt.Scanln(&choice)
	fmt.Println()

	fmt.Printf("You picked %v and computer picked %v\n", printAction(Action(choice)), printAction(Action(compChoice)))
	getWinner(Action(choice), Action(compChoice))
}
