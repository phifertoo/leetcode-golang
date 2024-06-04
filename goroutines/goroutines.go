package goroutines

import (
	"fmt"
	"time"
)

func printFunc(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func going(msg string) {
	fmt.Println(msg)
}

// even though the synchronous function is called after the asynchronous funtions, the synchronous function resolves first
func RoutineFunc() {
	go going("going")         // async function called
	go printFunc("goroutine") // async function called
	printFunc("direct")       // synchronous function called
	time.Sleep(time.Second)   // the RoutineFunc is called in main() which is a goroutine. Without this timer, the async functions would not finish before the main() function exits
	fmt.Println("done")
}
