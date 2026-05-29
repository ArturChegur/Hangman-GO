package main

import "fmt"

const (
	commandPlay = "play"
	commandExit = "exit"
)

func main() {
	err := LoadWords()

	if err != nil {
		fmt.Println("Failed to load words:", err)
		return
	}

	for {
		fmt.Print("Type [play] to play or [exit] to exit the game: ")
		var playerInput string

		_, err := fmt.Scan(&playerInput)
		if err != nil {
			return
		}

		switch playerInput {
		case commandPlay:
			PlayGame()
		case commandExit:
			return
		default:
			fmt.Println("Please choose [play] or [exit]")
		}
	}
}
