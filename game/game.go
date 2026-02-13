package game

// . is alias for importing without package name e.g.
// instead of io.Something() -> Something()
import (
	"bufio"
	"fmt"
	"time"
	. "wordle/io"
	. "wordle/model"
	. "wordle/structs"
)

// ANSI colors for the prints
const green = "\u001B[32m"
const yellow = "\u001B[33m"
const white = "\u001B[37m"
const reset = "\033[0m"

// Print string/character in certain color. ANSI color prefix + string + Color reset.
//
// Uses fmt.Print() so add a newline after using the function if necessary.
func PrintLetter(s string, ANSI_color string) {
	fmt.Print(ANSI_color + s + reset)
}

// Creates an array of letters ranging from A to Z.
//
// Returns byte array.
func InitLetterArray() []byte {

	arr := []byte{}

	for i := 0; i < 26; i++ {

		arr = append(arr, byte('A'+i))
	}

	return arr
}

// Takes lowercase character in bytes and substracts 32 to make it uppercase.
// If the byte value does not fit inbetween 'a' and 'z', the function will return the original character.
func ToUpperCase(char byte) byte {

	if char >= 'a' && char <= 'z' {
		return char - 32
	}
	return char
}

// Function will check if to_remove array items are included in the another array
// and removed in case they do.
//
// Returns an array which does not include the items.
func RemoveFromLetterArr(to_remove []byte, arr []byte) []byte {

	to_rm_len := len(to_remove)
	arr_len := len(arr)

	/*Outer loop is looping to remove chars*/
	for i := 0; i < to_rm_len; i++ {

		/*Inner loop for checking if the char is in arr*/
		for j := 0; j < arr_len; j++ {

			if ToUpperCase(to_remove[i]) == arr[j] {

				// New slice without the element and break out from inner loop
				arr = append(arr[:j], arr[j+1:arr_len]...)
				arr_len = len(arr)
				break
			}
		}

	}

	return arr
}

// Remove byte from array of strings if present
func RemoveCharFromArr(char byte, arr []string) []string {

	arr_len := len(arr)

	for i := 0; i < arr_len; i++ {

		if arr[i][0] == char {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

// Function checks if the char appears in the array
//
// if it does, the letter can be basically yellow
func CanBeYellow(char byte, arr []string) bool {

	arr_len := len(arr)

	for i := 0; i < arr_len; i++ {

		if arr[i][0] == char {

			return true
		}
	}

	return false
}

// Boolean value for byte appearing in string
func byteInString(b byte, s string) bool {

	for i := 0; i < 5; i++ {

		if b == s[i] {

			return true
		}
	}
	return false
}

// String "abc" to ["a", "b", "c"]
func StringToStringArr(s string) []string {
	s_len := len(s)
	arr := []string{}

	for i := 0; i < s_len; i++ {
		arr = append(arr, string(s[i]))
	}

	return arr
}

// Function parameters:
//
//	guess, word_to_guess string
//
// Each string has to be 5 characters.
//
// Funtion output:
//
//	[][]byte
//
// The function returns 2d byte array where each inner arrays contains a
// letter at index 0 and it's color at index 1.
// The colors are 'G', 'Y' or 'N'.
func GetFeedback(guess, word_to_guess string) (guess_letters [][]byte, chars_to_remove []byte) {

	non_green_letters := StringToStringArr(word_to_guess)
	guess_letters = [][]byte{}
	chars_to_remove = []byte{}
	// e.g. [[H, G], [E, Y], [L, G]]

	// Check green letters and fill array
	for i := 0; i < 5; i++ {

		color := byte('N')
		// Check if the letters are the same at i
		if guess[i] == word_to_guess[i] {

			color = byte('G')
			non_green_letters = RemoveCharFromArr(guess[i], non_green_letters)
		}

		guess_letters = append(guess_letters, []byte{guess[i], color})
	}
	// Check yellow and gray letters
	for i := 0; i < 5; i++ {

		char := guess_letters[i][0]
		color := guess_letters[i][1]

		// If there's an extra letter and there's an
		if color != byte('G') && CanBeYellow(char, non_green_letters) {

			guess_letters[i][1] = byte('Y')
			non_green_letters = RemoveCharFromArr(char, non_green_letters)
		} else if !byteInString(char, word_to_guess) {

			chars_to_remove = append(chars_to_remove, char)
		}

	}

	return
}

// The function will add the byte into the byte array IF it is an unique value.
//
// Returns a byte array where the item will be included.
func AddIfNotInArr(arr []byte, char byte) []byte {

	arr_len := len(arr)

	for i := 0; i < arr_len; i++ {

		if arr[i] == char {

			return arr
		}
	}

	arr = append(arr, char)

	return arr
}

// Used for actually printing the colored letters on the screen.
func PrintStats(feedback_arr [][]byte) {

	arr_len := len(feedback_arr)

	for i := 0; i < arr_len; i++ {

		color_letter := feedback_arr[i][1]
		guess_letter := ToUpperCase(feedback_arr[i][0])

		switch color_letter {
		case 'G':

			PrintLetter(string(guess_letter), green)
		case 'Y':

			PrintLetter(string(guess_letter), yellow)
		default:

			PrintLetter(string(guess_letter), white)
		}
	}
	fmt.Println()
}

// Function for the actual game. The function returns GameResultObject to main.
func Game(scanner *bufio.Scanner, WORDLISTPATH, word_to_quess, username string) GameResultObject {

	wordlist_offset_map := CreateWordListOffsetMap(WORDLISTPATH)
	letters_b_array := InitLetterArray()

	var attempts int
	var win bool = false
	var guessing_start int64 = time.Now().Unix()

	for attempts = 1; attempts < 7; attempts++ {

		guess := GetGuess(scanner, wordlist_offset_map, WORDLISTPATH)

		if word_to_quess == guess {
			fmt.Println("Congrats " + username + "! You've guessed the word correctly.")
			win = true
			break
		}

		// Give feedback
		feedback_arr, letters_to_rm := GetFeedback(guess, word_to_quess)
		PrintStats(feedback_arr)
		letters_b_array = RemoveFromLetterArr(letters_to_rm, letters_b_array)
		fmt.Println("Remaining letters: ", string(letters_b_array))
		fmt.Println()

	}

	if !win {
		fmt.Println("Game lost :( Better luck next time... The word was: " + word_to_quess)
	}

	var elapsed int64 = time.Now().Unix() - guessing_start

	game_stats := GameResultObject{}
	game_stats.User = username
	game_stats.SecretWord = word_to_quess
	game_stats.Attempts = attempts - 1
	game_stats.WinLose = win
	game_stats.Stime = guessing_start
	game_stats.Elapsed = elapsed

	return game_stats
}
