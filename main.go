package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	. "wordle/game"
	. "wordle/io"
	. "wordle/model"
	_ "wordle/structs"
)

func GetRandomNumber() int {
	// https://pkg.go.dev/math/rand#Rand.Intn
	// Simply return random int between 0 - 14 854
	return rand.Intn(14854)
}

func ValidateWordList() {
	// check file size is equal to the original

}

func main() {
	WORDLISTPATH, _ := filepath.Abs("./wordle-words.txt")
	CSVPATH, _ := filepath.Abs("./stats.csv")

	scanner := bufio.NewScanner(os.Stdin)
	username := GetUser(scanner)

	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")

	for {

		word_to_guess := string(ReadWordByRow(GetRandomNumber(), WORDLISTPATH))
		game_result_object := Game(scanner, WORDLISTPATH, word_to_guess, username)
		AppendToCSV(CSVPATH, game_result_object)

		if !StartNewGame(scanner) {
			break
		}

	}

	fmt.Println("See you again!")
}
