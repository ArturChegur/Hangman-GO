package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

var stdinReader = bufio.NewReader(os.Stdin)

func readLine(prompt string) (string, error) {
	fmt.Print(prompt)

	userInput, err := stdinReader.ReadString('\n')

	if err != nil {
		if !errors.Is(err, io.EOF) {
			return "", err
		}

		if strings.TrimSpace(userInput) == "" {
			return "", io.EOF
		}
	}

	return strings.TrimSpace(userInput), nil
}

func ScanCommand() (string, error) {
	for {
		userInput, err := readLine("Type [play] to play or [exit] to exit the game: ")
		if err != nil {
			return "", err
		}

		if userInput == "" {
			fmt.Println("Input cannot be empty!")
			continue
		}

		return strings.ToLower(userInput), nil
	}
}

func ScanSymbol() (rune, error) {
	for {
		userInput, err := readLine("Please enter [symbol]: ")
		if err != nil {
			return 0, err
		}

		runes := []rune(userInput)

		if len(runes) != 1 {
			fmt.Println("Input must be exactly 1 symbol!")
			continue
		}

		return unicode.ToLower(runes[0]), nil
	}
}
