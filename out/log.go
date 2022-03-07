package out

import "fmt"

func Info(message string) {
	fmt.Printf("\x1b[36m«info»\x1b[0m %s\n", message)
}

func Warn(message string) {
	fmt.Printf("\x1b[33m«warn»\x1b[0m %s\n", message)
}

func Err(message string) {
	fmt.Printf("\x1b[31m«err»\x1b[0m %s\n", message)
}
