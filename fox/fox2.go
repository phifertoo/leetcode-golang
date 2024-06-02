package fox

// element type
type Elem struct {
	Value     int
	Frequency int
}

// minHeap type
type Heap struct {
	Elements []Elem
}

func (h *Heap) insertAndSort(element Elem) {
	// append element
	h.Elements = append(h.Elements, element)
	// sort bottom up!!!
	currentIndex := len(h.Elements) - 1
	for currentIndex > 0 {
		parentIndex := (currentIndex - 1) / 2
		if h.Elements[parentIndex].Frequency < h.Elements[currentIndex].Frequency {
			break
		}
		// swap parent with child
		h.Elements[parentIndex], h.Elements[currentIndex] = h.Elements[currentIndex], h.Elements[parentIndex]
		// move up the tree
		currentIndex = parentIndex
	}
}

// func (h *Heap) sort() {
// 	currentIndex := len(h.Elements) - 1
// 	for currentIndex > 0 {
// 		parentIndex := (currentIndex - 1) / 2
// 		if h.Elements[parentIndex].Frequency < h.Elements[currentIndex].Frequency {
// 			break
// 		}
// 		// swap parent with child
// 		h.Elements[parentIndex], h.Elements[currentIndex] = h.Elements[currentIndex], h.Elements[parentIndex]
// 		// move up the tree
// 		currentIndex = parentIndex
// 	}
// }

func (h *Heap) removeMin() Elem {
	// root is always the minimum
	min := h.Elements[0]
	h.Elements[0] = h.Elements[len(h.Elements)-1]
	h.Elements = h.Elements[:len(h.Elements)-1]
	h.downHeap(0)

	return min
}

func (h *Heap) downHeap(parentIndex int) {
	lastIndex := len(h.Elements) - 1
	for {
		leftChildIndex := parentIndex*2 + 1
		rightChildIndex := parentIndex*2 + 2
		smallestIndex := parentIndex

		if leftChildIndex <= lastIndex && h.Elements[smallestIndex].Frequency > h.Elements[leftChildIndex].Frequency {
			smallestIndex = leftChildIndex
		}

		if rightChildIndex <= lastIndex && h.Elements[smallestIndex].Frequency > h.Elements[rightChildIndex].Frequency {
			smallestIndex = rightChildIndex
		}

		if smallestIndex == parentIndex {
			break
		}

		h.Elements[smallestIndex], h.Elements[parentIndex] = h.Elements[parentIndex], h.Elements[smallestIndex]
		parentIndex = smallestIndex
	}
}

func KMostFrequent(nums []int, k int) []int {
	// create frequencyMap
	freqMap := make(map[int]int)
	for _, num := range nums {
		_, exists := freqMap[num]
		if exists {
			freqMap[num]++
		} else {
			freqMap[num] = 1
		}
	}
	// initialize minHeap
	minHeap := Heap{}
	// iterate through frequencyMap to add element and sort to heap
	for num, frequency := range freqMap {
		element := Elem{
			Value:     num,
			Frequency: frequency,
		}
		minHeap.insertAndSort(element)
		// if len(heap) > k, remove the root then sort
		if len(minHeap.Elements) > k {
			minHeap.removeMin()
		}
	}
	// the root will then be out of order so you need to reorder
	// add elements from the heap to the output slice
	output := make([]int, k)
	for i := 0; i < k; i++ {
		output[k-i-1] = minHeap.removeMin().Value
	}

	return output
}
