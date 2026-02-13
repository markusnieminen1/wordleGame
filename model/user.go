package model

import (
	"bufio"
	"fmt"
	. "wordle/io"
)

func GetUser(scanner *bufio.Scanner) string {

	var input string

	for {
		fmt.Print("Enter your username: ")
		scanner.Scan()

		err := scanner.Err()

		if err != nil {
			fmt.Println("An error occured while reading username")
			continue
		}

		input = scanner.Text()

		if len(input) > 0 {
			break
		}

		fmt.Println("Not valid username. Try again...")
	}
	fmt.Print("\n")

	return input
}

func StartNewGame(scanner *bufio.Scanner) bool {

	var input string

	for {
		fmt.Print("Do you want to play again? (y/n): ")
		scanner.Scan()

		err := scanner.Err()

		if err != nil {
			fmt.Println("An error occured while reading user input...")
			continue
		}

		input = ToLower(scanner.Text())

		if len(input) != 1 {
			fmt.Println("Please enter valid response. y or n. ")
			continue
		}

		if input[0] == 'y' || input[0] == 'n' {
			break
		}

		fmt.Println("Not valid username. Try again...")
	}

	return input[0] == 'y'
}

func GetGuess(scanner *bufio.Scanner, offsetmap map[[2]byte][2]int, wordlistpath string) string {

	var input string

	for {
		fmt.Print("Enter your guess: ")
		scanner.Scan()

		err := scanner.Err()

		if err != nil {
			fmt.Println("An error occured while reading username")
			continue
		}

		input = scanner.Text()

		if len(input) != 5 {
			fmt.Println("Your guess must be exactly 5 letters long.")
			continue
		}

		if HasNonChars(input) {
			fmt.Println("Your guess must only contain lowercase letters.")
			continue
		}

		input = ToLower(input)

		offsetrange := GetOffset(input, offsetmap)

		if WordExists(wordlistpath, input, offsetrange) {
			break
		}

		fmt.Println("Word not in list. Please enter a valid word.")
	}

	fmt.Print("\n")
	return input

}

func ToUpper(s string) string {

	s_len := len(s)
	out := ""

	for i := 0; i < s_len; i++ {

		if s[i] >= 'a' && s[i] <= 'z' {

			out += string(s[i] - 32)
		} else {

			out += string(s[i])
		}
	}

	return out

}

func ToLower(s string) string {

	s_len := len(s)
	out := ""

	for i := 0; i < s_len; i++ {

		if s[i] >= 'A' && s[i] <= 'Z' {

			out += string(s[i] + 32)
		} else {

			out += string(s[i])
		}
	}

	return out
}

// Check if the string has other chars than letters
func HasNonChars(s string) bool {

	s_len := len(s)

	for i := 0; i < s_len; i++ {

		if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z') {

			continue
		} else {

			return true
		}
	}

	return false
}
