package calc

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

// TestAdd is a classic table-driven test: one table, one loop, many cases.
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"zeros", 0, 0, 0},
		{"positives", 2, 3, 5},
		{"negatives", -4, -6, -10},
		{"mixed", -2, 7, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // each case is a named subtest
			t.Parallel() // safe: since Go 1.22 tt is a fresh variable per iteration
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got, err := Divide(10, 2)
		assertNoError(t, err)
		if got != 5 {
			t.Fatalf("Divide(10, 2) = %d, want 5", got)
		}
	})
	t.Run("by zero", func(t *testing.T) {
		_, err := Divide(1, 0)
		if !errors.Is(err, ErrDivideByZero) {
			t.Fatalf("err = %v, want ErrDivideByZero", err)
		}
	})
}

// assertNoError is a helper; t.Helper makes failures point at the caller's line.
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestJoinVariantsAgree(t *testing.T) {
	words := []string{"go", "is", "fun"}
	if a, b := JoinSlow(words), JoinFast(words); a != b {
		t.Fatalf("JoinSlow=%q JoinFast=%q", a, b)
	}
}

// FuzzReverse lets the fuzzer hunt for inputs that break our invariants.
func FuzzReverse(f *testing.F) {
	for _, seed := range []string{"", "hello", "héllo", "日本語"} {
		f.Add(seed) // seed corpus; also runs as a normal test
	}
	f.Fuzz(func(t *testing.T, s string) {
		if !utf8.ValidString(s) {
			t.Skip("Reverse is defined for valid UTF-8 only")
		}
		rev := Reverse(s)
		if Reverse(rev) != s {
			t.Errorf("double reverse of %q = %q", s, Reverse(rev))
		}
		if utf8.RuneCountInString(rev) != utf8.RuneCountInString(s) {
			t.Errorf("rune count changed for %q", s)
		}
	})
}

var words = strings.Fields(strings.Repeat("gopher ", 1000))

// BenchmarkJoinFast uses the classic b.N loop, which still works everywhere.
func BenchmarkJoinFast(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JoinFast(words)
	}
}

// ExampleReverse is documentation AND a test: go test checks the Output comment.
func ExampleReverse() {
	fmt.Println(Reverse("stressed"))
	// Output: desserts
}
