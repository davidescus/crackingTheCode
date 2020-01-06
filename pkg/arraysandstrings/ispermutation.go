package arraysandstrings

import "sort"

// IsPermutationWithCounter will check if string s2 is a permutation of string s1
// by counting each unicode character
// time complexity: O(n)
// space complexity: O(1) (limited number of character which does not depend on strings length)
func IsPermutationWithCounter(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var occurrences = make(map[int32]int)
	for _, runeVal := range s1 {
		occurrences[runeVal]++
	}

	for _, runeVal := range s2 {
		occurrences[runeVal]--
		if occurrences[runeVal] < 0 {
			return false
		}
	}

	return true
}

// IsPermutationWithSort will check if string s2 is a permutation of string s1
// by sorts them and comparing
// time complexity: O(n)
// space complexity: O(n)
func IsPermutationWithSort(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	r1 := sortString(s1)
	r2 := sortString(s2)

	for i, v := range r1 {
		if v != r2[i] {
			return false
		}
	}
	return true
}

func sortString(s string) []rune {
	var r []rune
	for _, rv := range s {
		r = append(r, rv)
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return r
}
