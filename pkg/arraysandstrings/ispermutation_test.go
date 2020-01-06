package arraysandstrings

import "testing"

var resultIsPermutation bool

type testDataIsPermutation struct {
	name   string
	s1, s2 string
	exp    bool
}

func TestIsPermutationWithCounter(t *testing.T) {
	for _, tc := range getTestDataIsPermutation() {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPermutationWithCounter(tc.s1, tc.s2)
			if got != tc.exp {
				t.Errorf("got: %v but expect: %v for: '%s' vs '%s'", got, tc.exp, tc.s1, tc.s2)
			}
		})
	}
}

func TestIsPermutationWithSort(t *testing.T) {
	for _, tc := range getTestDataIsPermutation() {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPermutationWithSort(tc.s1, tc.s2)
			if got != tc.exp {
				t.Errorf("got: %v but expect: %v for: '%s' vs '%s'", got, tc.exp, tc.s1, tc.s2)
			}
		})
	}
}

func BenchmarkIsPermutationWithCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range getTestDataIsPermutation() {
			res := IsPermutationWithCounter(tc.s1, tc.s2)
			resultIsPermutation = res
		}
	}
}

func BenchmarkIsPermutationWithSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range getTestDataIsPermutation() {
			res := IsPermutationWithSort(tc.s1, tc.s2)
			resultIsPermutation = res
		}
	}
}

func getTestDataIsPermutation() []testDataIsPermutation {
	return []testDataIsPermutation{
		{
			name: "NotSameLength",
			s1:   "aaa",
			s2:   "aaaa",
			exp:  false,
		},
		{
			name: "ShortPermutation",
			s1:   "abc",
			s2:   "bca",
			exp:  true,
		},
		{
			name: "NotPermutation",
			s1:   "abc",
			s2:   "aaa",
			exp:  false,
		},
		{
			name: "SpecialCharsPermutation",
			s1:   "日本語語漢字かんじ",
			s2:   "日語本語じ漢字かん",
			exp:  true,
		},
		{
			name: "SpecialCharsNOTPermutation",
			s1:   "日本語語漢字かんじ",
			s2:   "日本語漢かんじ字字",
			exp:  false,
		},
		{
			name: "EmptyStringPermutation",
			s1:   "",
			s2:   "",
			exp:  true,
		},
	}
}
