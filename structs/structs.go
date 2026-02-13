package structs

// Object for csvs
// Anything starting with upper letter is visible to files outside this file
type GameResultObject struct {
	User       string
	SecretWord string
	Attempts   int
	WinLose    bool
	Stime      int64
	Elapsed    int64
}
