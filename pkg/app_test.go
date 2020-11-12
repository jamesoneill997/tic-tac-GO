package main

import (
	"testing"
)

//012 345 678 036 147 258 048 246
var testStrings = [][9]string{{"O", "O", "O", "3", "4", "5", "6", "7", "8"}, {"0", "1", "2", "O", "O", "O", "6", "7", "8"}, {"0", "1", "2", "3", "4", "5", "O", "O", "O"}, {"O", "1", "2", "O", "4", "5", "O", "7", "8"}, {"0", "O", "2", "3", "O", "5", "6", "O", "8"}, {"0", "1", "O", "3", "4", "O", "6", "7", "O"}, {"O", "1", "2", "3", "O", "5", "6", "7", "O"}, {"0", "1", "O", "3", "O", "5", "O", "7", "8"}}

func testCheckWinner(t *testing.T) {
	for str := range testStrings {
		result := CheckWinner(testStrings[str])
		if !result {
			t.Errorf("CheckWinner(%v) returned false, should have returned true.", testStrings[str])
		}
	}

}
