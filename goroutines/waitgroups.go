package goroutines

import (
	"fmt"
	"sync"
)

func runner1(wg *sync.WaitGroup) {
	defer wg.Done() // This decreases counter by 1
	fmt.Print("\nI am first runner")
}
func runner2(wg *sync.WaitGroup) {
	defer wg.Done() // This decreases counter by 1
	fmt.Print("\nI am second runner")
}
func Waitgroups() {
	wg := new(sync.WaitGroup)
	//increase waitgroup counter by 2 since we will run 2 goroutines
	wg.Add(2)
	go runner1(wg)
	go runner2(wg)
	// without wg.Wait() runner1 and runner2 would not print
	wg.Wait()
}
