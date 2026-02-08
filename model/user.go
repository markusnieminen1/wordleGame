package model

import (
	"fmt"
	"bufio"
	."wordle/io"
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

		fmt.Println("Not valid username. Try again...\n")
	}
	fmt.Print("\n")
	return input
}

func GetGuess(scanner *bufio.Scanner) string {

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

		off := GetOffset(guess, wordlist_offset_map)

		if WordExists(WORDLISTPATH, guess, off) {
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
		if s[i] >= 'a' && s[i] < = 'z' {
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
		if s[i] >= 'A' && s[i] < = 'Z' {
			out += string(s[i] + 32)
		} else {
			out += string(s[i])
		}
	}
	return out 

}

func HasNonChars(s string) bool {

	s_len := len(s)

	for i := 0; i < s_len; i++ {
		if (s[i] >= 'A' && s[i] < = 'Z') || (s[i] >= 'a' && s[i] < = 'z'){
			continue
		} else {
			return true
		}
	}
	
	return false 

}