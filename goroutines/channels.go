package goroutines

import "fmt"

func Channels() {
	// create a channel called "messages" that moves a string
	messages := make(chan string)

	go sendPing(messages)
	output := receivePing(messages)
	close(messages)
	// If you call receivePing directly in the main function or any other blocking (synchronous) part of your code, it will wait (block) until a message is available on the channel.
	fmt.Print(output)
}

func sendPing(ch chan<- string) {
	ch <- "ping"
}

func receivePing(ch <-chan string) string {
	msg := <-ch

	return msg
}
