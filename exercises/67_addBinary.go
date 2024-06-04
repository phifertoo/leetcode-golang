package exercises

import (
	"fmt"
)

// Given two binary strings a and b, return their sum as a binary string.
func AddBinary(a string, b string) string {
	// determine which string is shorter
	shorter := a
	longer := b
	if len(b) < len(a) {
		shorter = b
		longer = a
	}
	// carry one
	carryOne := false
	// iterate backwards on shorter string

	output := ""
	for i := len(shorter) - 1; i >= 0; i-- {
		longerIndex := len(longer) - 1 - (len(shorter) - 1 - i)
		if carryOne {
			if shorter[i] == '0' && longer[longerIndex] == '0' {
				carryOne = false
				output = string('1') + output
			} else {
				if shorter[i] == '1' && longer[longerIndex] == '1' {
					output = string('1') + output
				} else {
					// shorter == 1 && longer == 1
					output = string('0') + output
				}

				carryOne = true
			}
		} else {
			if shorter[i] == '1' && longer[longerIndex] == '1' {
				carryOne = true
				output = string('0') + output
			} else if shorter[i] == '0' && longer[longerIndex] == '0' {
				carryOne = false
				output = string('0') + output
			} else {
				carryOne = false
				output = string('1') + output
			}
		}
	}

	// if carryOne is still true at the end of the loop, we need to add 1 to the "end" of longer string then prepend the rest of the longer string
	// add the remainer of the longer to the string
	for i := len(longer) - len(shorter) - 1; i >= 0; i-- {
		if carryOne {
			if longer[i] == '1' {
				carryOne = true
				output = string('0') + output
			}
			if longer[i] == '0' {
				carryOne = false
				output = string('1') + output
			}

		} else {
			output = string(longer[i]) + output
		}
	}

	// after iterating through the longer string, carryOne may still be true so we need to prepend another "1"
	if carryOne {
		output = string("1") + output
	}

	return output
}

// func AddBinary(a string, b string) string {
// 	// determine longer string
// 	longer := ""
// 	shorter := ""
// 	if len(a) >= len(b) {
// 		longer = a
// 		shorter = b
// 	} else {
// 		longer = b
// 		shorter = a
// 	}

// 	// create output slice
// 	output := []string{}
// 	for i := 0; i < len(longer); i++ {
// 		output = append(output, string(longer[i]))
// 	}

// 	carryOver := false
// 	for i := len(longer) - 1; i >= 0; i-- {
// 		shortIndex := len(shorter) - 1 - (len(longer) - 1 - i)
// 		shortNum := "0"
// 		if shortIndex > -1 {
// 			shortNum = string(shorter[shortIndex])
// 		} else {
// 			shortNum = "0"
// 		}

// 		num, next := handleBinary(shortNum, string(longer[i]), carryOver)
// 		if next {
// 			carryOver = true
// 		} else {
// 			carryOver = false
// 		}

// 		output[i] = num
// 	}

// 	if carryOver {
// 		output = append([]string{"1"}, output...)
// 	}

// 	return strings.Join(output, "")
// }

// func handleBinary(numA string, numB string, carryOver bool) (string, bool) {
// 	if carryOver {
// 		if numA == "1" && numB == "1" {
// 			return "1", true
// 		}

// 		if numA == "1" || numB == "1" {
// 			return "0", true
// 		}

// 		return "1", false
// 	}

// 	if numA == "1" && numB == "1" {
// 		return "0", true
// 	}

// 	if numA == "1" || numB == "1" {
// 		return "1", false
// 	}

// 	return "0", false
// }

func AddBinaryTester() bool {
	fmt.Print(AddBinary("11", "1"))      // "100"
	fmt.Print(AddBinary("1010", "1011")) //"10101"

	return true
}
