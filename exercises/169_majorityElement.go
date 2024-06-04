package exercises

import "fmt"

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	frequencyMap := map[int]int{}
	for _, num := range nums {
		_, ok := frequencyMap[num]
		if ok {
			frequencyMap[num]++
			if frequencyMap[num] > len(nums)/2 {
				return num
			}
		} else {
			frequencyMap[num] = 1
		}
	}
	return 0
}

// func majorityElement(nums []int) int {
// 	numberMap := make(map[int]int)

// 	for _, num := range nums {
// 		_, exists := numberMap[num]
// 		if exists {
// 			numberMap[num]++
// 		} else {
// 			numberMap[num] = 1
// 		}

// 		if numberMap[num] > (len(nums) / 2) {
// 			return num
// 		}
// 	}

// 	// fmt.Printf("%+v\n", numberMap)

// 	return 0
// }

func MajorityElementTester() bool {
	fmt.Print(majorityElement([]int{3, 2, 3}))             // 3
	fmt.Print(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // 2

	return true
}
