// Errors: values, wrapping with %w, errors.Is / errors.As, Join, panic/recover.
//
// Run from assets/code:
//
//	go run ./ch03/errors
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

// ErrNotFound is a sentinel error: a package-level value callers compare against.
var ErrNotFound = errors.New("not found")

// ValidationError is a custom error type carrying structured data.
type ValidationError struct {
	Field string
	Value string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s: %q", e.Field, e.Value)
}

var users = map[int]string{1: "ada", 2: "grace"}

func findUser(id int) (string, error) {
	name, ok := users[id]
	if !ok {
		return "", fmt.Errorf("findUser(%d): %w", id, ErrNotFound) // wrap, keep the chain
	}
	return name, nil
}

func parseAge(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil || n < 0 {
		return 0, &ValidationError{Field: "age", Value: s}
	}
	return n, nil
}

func handler(id int) error {
	if _, err := findUser(id); err != nil {
		return fmt.Errorf("handler: %w", err) // add context at each layer
	}
	return nil
}

// safeDivide turns a panic into an ordinary error at an API boundary.
func safeDivide(a, b int) (q int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	return a / b, nil
}

func main() {
	err := handler(42)
	fmt.Println(err)
	fmt.Println("is ErrNotFound?", errors.Is(err, ErrNotFound))

	_, err = parseAge("-3")
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("bad field:", ve.Field)
	}

	_, err = os.Open("does-not-exist.txt")
	fmt.Println("is fs.ErrNotExist?", errors.Is(err, fs.ErrNotExist))

	joined := errors.Join(ErrNotFound, &ValidationError{"email", "nope"})
	fmt.Println("joined:", joined)
	fmt.Println("joined is ErrNotFound?", errors.Is(joined, ErrNotFound))

	_, err = safeDivide(1, 0)
	fmt.Println(err)
}
