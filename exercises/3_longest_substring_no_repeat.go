package exercises

// Given a string s, find the length of the longest
// substring
//  without repeating characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	window := []string{}
	// sliding window
	// keep appending if not found in map of window
	// if a match is found, we start the new window at the index of the duplicate
	// abcdd
	for _, letter := range s {
		// letterExistsInWindow := false
		indexFound := -1
		for i, l := range window {
			if l == string(letter) {
				indexFound = i
				break
			}
		}

		if indexFound > -1 {
			// reset window to new index
			window = window[indexFound+1:]
			window = append(window, string(letter))
		} else {
			window = append(window, string(letter))
			if len(window) > maxLength {
				maxLength = len(window)
			}
		}
	}
	return maxLength
}

func LengthOfLongestSubstringTester() bool {
	fmt.Print(lengthOfLongestSubstring("abcbbabcde")) // 3

	return true
}
