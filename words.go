package main

import (
	"math/rand"
	"os"
	"strings"
)

const wordsFileName = "words.txt"

var words []string

func LoadWords() error {
	data, err := os.ReadFile(wordsFileName)

	if err != nil {
		return err
	}

	words = strings.Fields(string(data))

	return nil
}

func GetRandomWord() string {
	if len(words) == 0 {
		return ""
	}

	return words[rand.Intn(len(words))]
}
