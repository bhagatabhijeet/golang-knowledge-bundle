// Package greet builds friendly greetings. It exists so the Go CLI lessons
// have a real package to format, test and import.
package greet

import "strings"

// Hello returns a greeting for name. An empty name falls back to "World".
// When shout is true the greeting is upper-cased.
func Hello(name string, shout bool) string {
	if strings.TrimSpace(name) == "" {
		name = "World"
	}
	msg := "Hello, " + name + "!"
	if shout {
		msg = strings.ToUpper(msg)
	}
	return msg
}
