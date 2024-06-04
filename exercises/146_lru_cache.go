package exercises

import (
	"container/list"
)

// LRUCache is a struct representing an LRU (Least Recently Used) cache
type LRUCache struct {
	capacity int                   // Maximum capacity of the cache
	cache    map[int]*list.Element // Map to store key-value pairs, where value is a pointer to an element in the doubly linked list
	order    *list.List            // Doubly linked list to maintain the order of elements
}

// entry is a struct representing a key-value pair stored in the LRUCache
type entry struct {
	key, value int
}

// Constructor initializes and returns a new LRUCache with the given capacity
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		order:    list.New(),
	}
}

// Get retrieves the value associated with the given key from the cache
// If the key is found, it moves the corresponding element to the front of the order list (most recently used)
// If the key is not found, it returns -1
func (lru *LRUCache) Get(key int) int {
	if element, ok := lru.cache[key]; ok { // Check if the key exists in the cache
		lru.order.MoveToFront(element)     // Move the element to the front (most recently used)
		return element.Value.(entry).value // Return the value associated with the key
	}
	return -1 // Key not found, return -1
}

// Put adds a new key-value pair to the cache, or updates the value if the key already exists
// If the cache exceeds its capacity, it evicts the least recently used element
func (lru *LRUCache) Put(key, value int) {
	if element, ok := lru.cache[key]; ok { // Check if the key already exists in the cache
		lru.order.MoveToFront(element)    // Move the element to the front (most recently used)
		element.Value = entry{key, value} // Update the value associated with the key
	} else {
		if lru.order.Len() == lru.capacity { // Check if the cache has reached its capacity
			backElement := lru.order.Back() // Get the least recently used element (back of the list)
			if backElement != nil {
				lru.order.Remove(backElement)                    // Remove the least recently used element from the list
				delete(lru.cache, backElement.Value.(entry).key) // Remove the corresponding entry from the cache
			}
		}
		newElement := lru.order.PushFront(entry{key, value}) // Add the new key-value pair to the front of the list
		lru.cache[key] = newElement                          // Add the new entry to the cache map
	}
}
