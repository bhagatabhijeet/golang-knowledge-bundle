package greet

import (
	"fmt"
	"testing"
)

// TestHello is a table-driven test: one row per case, one loop to run them all.
func TestHello(t *testing.T) {
	tests := []struct {
		name  string
		input string
		shout bool
		want  string
	}{
		{"plain", "Gopher", false, "Hello, Gopher!"},
		{"shout", "Gopher", true, "HELLO, GOPHER!"},
		{"empty name", "", false, "Hello, World!"},
		{"spaces only", "   ", false, "Hello, World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.input, tt.shout); got != tt.want {
				t.Errorf("Hello(%q, %v) = %q, want %q", tt.input, tt.shout, got, tt.want)
			}
		})
	}
}

// ExampleHello is documentation that go test also runs and checks.
func ExampleHello() {
	fmt.Println(Hello("Ada", false))
	// Output: Hello, Ada!
}
