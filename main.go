package main 

import (
	 "fmt"
	"path/filepath"
	"math/rand"
	"bufio"
	"os"
	_ "wordle/structs"
	. "wordle/game"
	_ "wordle/io"
	. "wordle/model"
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
	//CSVPATH, _ := filepath.Abs("./stats.csv")

	scanner := bufio.NewScanner(os.Stdin)
	username := GetUser(scanner)

	Game(scanner, WORDLISTPATH, "aalie")

	fmt.Print(username)



}