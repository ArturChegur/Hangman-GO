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
		playerInput, err := ScanCommand()

		if err != nil {
			fmt.Println("\nBye!")
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
