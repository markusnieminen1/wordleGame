package game

// . is alias for importing without package name e.g. instead of io.Something() -> Something()

import (
	. "wordle/io"
	. "wordle/model"
	"fmt"
	"bufio"
	)

// ANSI colors for the prints 
const green = "\u001B[32m"
const yellow = "\u001B[33m"
const white = "\u001B[37m"
const reset = "\033[0m"


func PrintLetter(s string, ANSI_color string) {
	fmt.Print(ANSI_color + s + reset)
}

func InitLetterArray() []string {
	
	arr := []string{}
	for i := 0; i < 26; i++ {
		arr = append(arr, string('A' + i))
	}

	return arr 
}


func RemoveFromLetterArr(to_remove string, arr []string) []string {

	word_len := len(to_remove)
	arr_len := len(arr)

	for i:=0; i < word_len; i++ {

		for j := 0; j < arr_len; j++ {
			
			if to_remove[i] == arr[j] && j + i < arr_len {
				arr = append(arr, [0:j], [j+1:arr_len-1]...)
			} else if to_remove[i] == arr[j] {
				arr = append(arr, [0:arr_len-2])
			}
		}

	}

	return arr
}

func PrintFeedback(guess, word_to_guess string){

	// Check corrects 
	arr := make([]string, 5)
	for i := 0; i < 5; i++ {
		if guess[i] == word_to_guess[i] {
			arr[i] = "Y"
		}
	}
	
	guess_map := make(map[byte]int)
	gu_len = len(s)

	for i := 0; i < gu_len-1; i++{
		if guess_map[guess[i]] {
			continue
		}
		count := 0
		for k := + 1; k < 5; k++ {
			if guess[k] == guess[i] {
				count += 1
			}
		}
		guess_map[guess[i]] = count
	}



}


func PrintStats() {

}

func Game(scanner *bufio.Scanner, WORDLISTPATH, word_to_quess string) {
	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")

	wordlist_offset_map := CreateWordListOffsetMap(WORDLISTPATH)
	letters_array := InitLetterArray() 
	var attempts int 
	// Guesses 
	for i := 1; i < 7; i++ {
	
		guess := GetGuess(scanner)

		if word_to_quess == guess {
			fmt.Println("Congratulations! You've guessed the word correctly.")
			attempts = i 
			break
		}
			
	}
	// stats 
	
	fmt.Println(exi, guess)
}