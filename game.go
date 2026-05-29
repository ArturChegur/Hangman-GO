package main

import "fmt"

type Game struct {
	word          []rune
	visibleWord   []rune
	usedLetters   []rune
	wrongAttempts int
}

func NewGame(word string) *Game {
	wordRunes := []rune(word)
	visibleWord := make([]rune, len(wordRunes))

	for i := range wordRunes {
		visibleWord[i] = '_'
	}

	return &Game{
		word:        wordRunes,
		visibleWord: visibleWord,
		usedLetters: make([]rune, 0, len(wordRunes)),
	}
}

func PlayGame() {
	randomWord, err := GetRandomWord()
	if err != nil {
		fmt.Println("Cannot start game:", err)
		return
	}

	game := NewGame(randomWord)

	for game.wrongAttempts < hangmanAttempts {
		drawHangman(game.wrongAttempts)
		game.printUsedLetters()
		game.printVisibleWord()
		fmt.Printf("Attempts left: %d\n", hangmanAttempts-game.wrongAttempts)

		symbol, err := ScanSymbol()

		if err != nil {
			fmt.Println("Error while input")
			return
		}

		if game.isUsedLetter(symbol) {
			fmt.Println("You already tried this letter!")
			continue
		}

		game.usedLetters = append(game.usedLetters, symbol)

		if game.isSymbolInWord(symbol) {
			game.changeVisibleWord(symbol)

			if game.isWin() {
				drawHangman(game.wrongAttempts)
				game.printVisibleWord()
				fmt.Printf("You won! The word was: %s\n", string(game.word))
				return
			}

			continue
		}

		game.wrongAttempts++
	}

	drawHangman(hangmanAttempts)
	fmt.Printf("You lost! The word was: %s\n", string(game.word))
}

func (game *Game) isUsedLetter(symbol rune) bool {
	for _, letter := range game.usedLetters {
		if letter == symbol {
			return true
		}
	}

	return false
}

func (game *Game) isSymbolInWord(symbol rune) bool {
	for _, letter := range game.word {
		if letter == symbol {
			return true
		}
	}

	return false
}

func (game *Game) changeVisibleWord(symbol rune) {
	for i, letter := range game.word {
		if letter == symbol {
			game.visibleWord[i] = symbol
		}
	}
}

func (game *Game) isWin() bool {
	for i, letter := range game.word {
		if game.visibleWord[i] != letter {
			return false
		}
	}

	return true
}

func (game *Game) printUsedLetters() {
	fmt.Printf("Already used letters: %s\n", string(game.usedLetters))
}

func (game *Game) printVisibleWord() {
	fmt.Printf("Current word: %s\n", string(game.visibleWord))
}
