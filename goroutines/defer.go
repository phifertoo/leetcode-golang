package goroutines

import "fmt"

func DeferFunc() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
