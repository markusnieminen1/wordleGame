package io

import (
	"os"
	"fmt"
	. "wordle/structs"
)
// 14855 rows times 6 bytes (5 chars, and linefeed)
const FILE_SIZE_IN_BYTES = 89130 

// General function for error handling
func VerifyAllesGut(err error) {
	if err != nil {
		fmt.Println("\033[31m\n\nAn Error Occured:")
		fmt.Println(err, "\033[0m")
		fmt.Println("\n\nShutting down...")
		os.Exit(1)
	}
}


// Read just the correct line of text and return byte array containing the word
func ReadWordByRow(line_number int, fullpath string) []byte {

	// All the words in the list contains 5 characters 
	// so we can use byte index to access the word we are
	// interested in
	// https://pkg.go.dev/os#File.ReadAt

	read_bytes_from := int64(6 * line_number)
	byte_arr := make([]byte, 5)

	// Open and close the file after the block. 
	// if err raises an error, it will be PathError 
	file, err := os.Open(fullpath)
	defer file.Close()

	if err != nil {
		fmt.Println("\033[31mPATH ERROR WHEN READING THE WORDS. CHECK THE WORD FILE EXISTS IN ROOT. \033[0m")
		os.Exit(1)
	}

	// return pos 0 is for bytes read. err non nill or io.EOF
	// Use = instead of := as the err is already declared and _ does not save the return var.
	_ , err = file.ReadAt(byte_arr, read_bytes_from)

	if err != nil {
		fmt.Println("\033[31mEND OF FILE OR NON NILL ERR. INDEX 14854 IS THE LAST POSSIBLE WORD. \033[0m")
	}

	return byte_arr
}

// Integer to string conversion
func IntToStr(in_int int) string {

	// Loop the numbers and add modulus of 10
	// Divide original number each time by 10
	num_arr := []int{}
	for in_int != 0 {

		num_arr = append(num_arr, in_int % 10)
		in_int = in_int / 10
	}

	// The list is reversed from the previous step, so start from the end
	// and change the values from int to text at the same time 
	text_arr := []byte{}
	for i := len(num_arr)-1; i >= 0; i-- {

		text_arr = append(text_arr, byte(48 + num_arr[i])) 
	}

	// Convert byte arr into string using string casting 
	return string(text_arr)
}


// String to byte array conversion
func StrToByteArr(s string) []byte {

	s_len := len(s)
	arr := []byte{}
	
	for i := 0; i < s_len; i++ {
		arr = append(arr, s[i])
	}

	return arr
}


// Function for appending rows to the CSV
func AppendToCSV(fullpath string, Obj GameResultObject) {

	// https://pkg.go.dev/os#OpenFile
	// https://pkg.go.dev/os#File.Stat
	// https://pkg.go.dev/io/fs#FileInfo

	// Append / CREATE to have less space for error
	file, err := os.OpenFile(fullpath, os.O_APPEND|os.O_WRONLY, 0644)
	VerifyAllesGut(err)
	defer file.Close()

	fileinfo, err := file.Stat()
	VerifyAllesGut(err)

	/* CSV FIELDS
    username
    secret word
    number of attempts
    "win" or "loss"
	*/
	

	// Check if the file needs columns before the actual data
	if fileinfo.Size() < 1 {
		// use `comments to include ""`
		file_columns := `username,secret word,number of attempts,"win" or "loss"`
		_, err = file.Write(StrToByteArr(file_columns))
		VerifyAllesGut(err)
	}
	
	// Create the object to the file 
	// Multiline function calls are not supported?
	comma := ","
	byte_arr := StrToByteArr(comma + string(10) + Obj.User + comma + Obj.Scrt + comma + Obj.C_att + comma + Obj.WL)
	_, err = file.Write(byte_arr)
	VerifyAllesGut(err)
	

	return
}

func CreateWordListOffsetMap(fullpath string) map[[2]byte][2]int {

	file, err := os.OpenFile(fullpath, os.O_RDONLY, 0644)
	VerifyAllesGut(err)
	defer file.Close()

	/*
	Find as high Int as possible that resolves 89130 % (x * 6) 
	ideally to 0 to have optimal amount of reads
	Memory usage being the trade off 
	Count of reads (blocking sys calls) versus memory usage 
	*/
	rows_at_a_time := 550
	chunk_size := 6 * rows_at_a_time

	byte_arr := make([]byte, chunk_size)

	offset_map := make(map[[2]byte][2]int)

	// starting with 'a' 'a' 
	var cur_chars_bytes [2]byte = [2]byte{97, 97}

	// For inner loop. create a pointer
	var i_byte_arr [2]byte = [2]byte{0, 0}
	var value_holder [2]int = [2]int{0, 0}
	ptr := &i_byte_arr 
	ptr2 := &value_holder

	var cur_pos int = 0

	/* ---------------------
	
	LOOP FOR THE FILE CHUNKS
	-----------------------*/
	for i := 0; i < FILE_SIZE_IN_BYTES; i += chunk_size {

		// Use whence: 0 -> relative to the file 
		// func (f *File) Seek(offset int64, whence int) (ret int64, err error)
		_, err = file.Seek(int64(i), 0)
		VerifyAllesGut(err)

		// Check if the file still has contents for "full" chunk read
		// and if not, just recreate the array
		if FILE_SIZE_IN_BYTES - i < chunk_size {
			byte_arr = make([]byte, FILE_SIZE_IN_BYTES % chunk_size)
		}

		_, err = file.Read(byte_arr)
		VerifyAllesGut(err)
		
		var arr_len int = len(byte_arr)

		/* ---------------------
	
		LOOP THE BYTES IN FILE CHUNKS 
		-----------------------*/
		for inner_i := 0; inner_i < arr_len; inner_i += 6 {
			
			// i + 1 will be always inside index, i + 7 would fail on last item
			ptr[0], ptr[1] = byte_arr[inner_i], byte_arr[inner_i + 1]

			if cur_chars_bytes != *ptr {
				ptr2[0], ptr2[1] = cur_pos, i + inner_i
				//offset_map[cur_chars_bytes] = []int{cur_pos, i + inner_i}
				offset_map[cur_chars_bytes] = *ptr2
				
				cur_chars_bytes[0], cur_chars_bytes[1] = byte_arr[inner_i], byte_arr[inner_i + 1]
				cur_pos = i + inner_i
			}

		}

	}

	// Add last row chars
	rest := make([]byte, 2)
	cur_chars_bytes = [2]byte{}

	_, err = file.Seek(int64(FILE_SIZE_IN_BYTES - 12), 0)
	VerifyAllesGut(err)
	
	_, err = file.Read(rest)
	VerifyAllesGut(err)

	copy(cur_chars_bytes[0:2], rest[0:2])

	offset_map[cur_chars_bytes] = [2]int{cur_pos, FILE_SIZE_IN_BYTES}

	return offset_map
}	


func ReadOffset(fullpath string, arr [2]int) []byte {

	file, err := os.Open(fullpath)
	defer file.Close()

	output_arr := make([]byte, arr[1] - arr[0])

	// Go to the location
	_, err = file.Seek(int64(arr[0]), 0)
	VerifyAllesGut(err)

	_, err = file.Read(output_arr)
	VerifyAllesGut(err)

	return output_arr

}

func GetOffset(word string, offsetmap map[[2]byte][2]int) [2]int {
	// Just return an element based on the words first 2 chars
	return offsetmap[[2]byte{word[0], word[1]}]

}


func WordExists(fullpath, guessed_word string, offset [2]int) bool {
	
	word_byte_arr := []byte{}
	for i := 0; i < 5; i++ {
		word_byte_arr = append(word_byte_arr, guessed_word[i])
	}

	byte_arr := ReadOffset(fullpath, offset)


	arr_len := len(byte_arr)
	for i := 0 ; i < arr_len; i += 6 {
		fmt.Println(string(byte_arr[i:i+5]), guessed_word)
		if guessed_word == string(byte_arr[i:i+5]) {
			return true 
		}
	}

	return false 

}

