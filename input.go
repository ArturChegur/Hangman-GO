package main

import "fmt"

func ScanSymbol() string {
	var userInput string

	for {
		_, err := fmt.Scan(&userInput)
		if err != nil {
			return ""
		}

		if len([]rune(userInput)) != 1 {
			fmt.Println("Input must be 1 symbol!")
			continue
		}

		return userInput
	}
}
