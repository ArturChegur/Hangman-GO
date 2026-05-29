package main

import (
	"errors"
	"math/rand"
	"os"
	"strings"
	"time"
)

const wordsFileName = "words.txt"

var words []string

var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))

func LoadWords() error {
	data, err := os.ReadFile(wordsFileName)

	if err != nil {
		return err
	}

	words = strings.Fields(string(data))
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	if len(words) == 0 {
		return errors.New("word list is empty")
	}

	return nil
}

func GetRandomWord() (string, error) {
	if len(words) == 0 {
		return "", errors.New(" word list is empty, call LoadWords first")
	}

	return words[randomizer.Intn(len(words))], nil
}
