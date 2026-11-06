// Command strings shows bytes vs runes and UTF-8 in practice.
//
//	go run ./ch02/strings
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "héllo, 世界"

	fmt.Println("bytes:", len(s))
	fmt.Println("runes:", utf8.RuneCountInString(s))

	// Indexing gives a byte, not a character.
	fmt.Printf("s[1] = %d (0x%x)\n", s[1], s[1])

	// range decodes UTF-8 and yields (byte offset, rune).
	for i, r := range "é世" {
		fmt.Printf("offset %d: %c U+%04X (%d bytes)\n", i, r, r, utf8.RuneLen(r))
	}

	// Convert to []rune for character-level slicing.
	r := []rune(s)
	fmt.Println(string(r[7:]))

	// Strings are immutable; build big ones with strings.Builder.
	var b strings.Builder
	for i := range 3 {
		fmt.Fprintf(&b, "[%d]", i)
	}
	fmt.Println(b.String())

	fmt.Println(strings.ToUpper("go"), strings.Fields(" a  b c "), strings.Repeat("=", 5))

	// Raw string literals keep backslashes and newlines as-is.
	raw := `C:\path\no\escapes`
	fmt.Println(raw)
}
