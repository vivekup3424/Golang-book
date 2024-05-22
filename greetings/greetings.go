package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %v.\n Welcome!\n", name)
	return message
}
