package exercises

import "fmt"

// You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

// Increment the large integer by one and return the resulting array of digits.

func PlusOne(digits []int) []int {
	shouldIncrement := true
	for i := len(digits) - 1; i >= 0; i-- {
		if shouldIncrement {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				shouldIncrement = false
				digits[i]++
				break
			}
		}
	}

	// this means the for loop ended and needs to prepend a 1 (ie 999)
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

// func PlusOne(digits []int) []int {

// 	incrementNext := false
// 	// needsIncrement := false
// 	for i := len(digits) - 1; i >= 0; i-- {
// 		// if i == len(digits)-1 {
// 		num, shouldIncrementNext := addOne(digits[i])
// 		incrementNext = shouldIncrementNext
// 		digits[i] = num
// 		if !shouldIncrementNext {
// 			incrementNext = false
// 			break
// 		}
// 	}

// 	// handle [9] => [1, 0], [9, 9, 9] => [1, 0, 0. 0]
// 	if incrementNext {
// 		digits = append([]int{1}, digits...)
// 	}

// 	return digits
// }

// func addOne(num int) (int, bool) {
// 	if num == 9 {
// 		return 0, true
// 	}

// 	return num + 1, false
// }

func PlusOneTester() bool {
	fmt.Print(PlusOne([]int{1, 2, 3}))    // [1,2,4]
	fmt.Print(PlusOne([]int{1, 2, 9}))    // [1,2,4]
	fmt.Print(PlusOne([]int{4, 3, 2, 1})) // [4,3,2,2]
	fmt.Print(PlusOne([]int{9}))          // [1,0]

	return true
}
