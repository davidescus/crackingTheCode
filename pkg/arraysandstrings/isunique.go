package arraysandstrings

import (
	"fmt"
	"unicode/utf8"
)

// IsUniqueCharactersIterative verify if a string contain only unique characters encoded in utf-8
// return err if could not decode character
// time complexity: O(n2)
// space complexity: O(1)
func IsUniqueCharactersIterative(s string) (bool, error) {
	for i, outer := range s {
		runeLen := utf8.RuneLen(outer)
		// \uFFFD represent unicode replacement character
		if runeLen < 0 || outer == '\uFFFD' {
			return false, fmt.Errorf("invalid utf-8 string at position: %d", i)
		}
		for _, inner := range s[i+runeLen:] {
			if outer == inner {
				return false, nil
			}
		}
	}
	return true, nil
}

// IsUniqueCharactersMemorize verify if a string contain only unique characters encoded in utf-8
// return err if could not decode character
// time complexity: O(n)
// space complexity: O(n)
func IsUniqueCharactersMemorize(s string) (bool, error) {
	mem := make(map[rune]bool)
	for i, outer := range s {
		// \uFFFD represent unicode replacement character
		if outer == '\uFFFD' {
			return false, fmt.Errorf("invalid utf-8 string at position: %d", i)
		}
		if exists := mem[outer]; exists {
			return false, nil
		}
		mem[outer] = true
	}
	return true, nil
}
