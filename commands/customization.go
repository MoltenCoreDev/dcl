package commands

import "fmt"

// TODO: Document this
func DrawPrompt(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
