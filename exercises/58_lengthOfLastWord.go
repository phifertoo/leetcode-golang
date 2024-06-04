package exercises

import (
	"fmt"
)

// Given a string s consisting of words and spaces, return the length of the last word in the string.

// A word is a maximal substring consisting of non-space characters only.
func lengthOfLastWord(s string) int {
	// find last space
	// reverse string
	reversedString := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversedString += string(s[i])
	}

	// find first space
	firstLetterIndex := -1
	for i, character := range reversedString {
		if string(character) != " " && firstLetterIndex == -1 {
			firstLetterIndex = i
		}

		if string(character) == " " && firstLetterIndex > -1 {
			// return oppositve
			return i - firstLetterIndex
		}
	}

	// return slice
	return len(s) - firstLetterIndex
}

// func LengthOfLastWord(s string) int {
// 	// find the last space
// 	// the string might end in a space
// 	// the string might not have a space
// 	startIndex := 0
// 	endIndex := 0
// 	startIndexFound := false
// 	endIndexFound := false
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if !endIndexFound && s[i] != ' ' {
// 			endIndexFound = true
// 			endIndex = i + 1
// 		}

// 		if !startIndexFound && endIndexFound && s[i] == ' ' {
// 			startIndexFound = true
// 			startIndex = i + 1
// 			break
// 		}
// 	}

// 	if !endIndexFound {
// 		return len(s)
// 	}

// 	return endIndex - startIndex
// }

func LengthOfLastWordTester() bool {
	fmt.Print(lengthOfLastWord("Hello World"))                 // 5
	fmt.Print(lengthOfLastWord("   fly me   to   the moon  ")) // 4
	fmt.Print(lengthOfLastWord("luffy is still joyboy"))       // 6
	fmt.Print(lengthOfLastWord("a "))                          // 1

	return true
}
