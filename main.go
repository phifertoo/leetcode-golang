package main

import (
	"fmt"

	"github.com/phifertoo/leetcode-golang/fox"
)

func main() {
	fmt.Printf("hello")
	heap := fox.MinHeap{}

	fmt.Print(heap.HighestFreq([]int{1, 1, 1, 2, 2, 3, 4, 4, 4}, 3))

}
