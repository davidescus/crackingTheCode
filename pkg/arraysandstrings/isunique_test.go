package arraysandstrings

import "testing"

var resultIsUnique bool

type testDataIsUniqueCharacters struct {
	Name   string
	Val    string
	Exp    bool
	HasErr bool
}

func TestIsUniqueCharactersIterative(t *testing.T) {
	for _, tc := range getTestDataIsUniqueCharacters() {
		t.Run(tc.Name, func(t *testing.T) {
			got, err := IsUniqueCharactersIterative(tc.Val)
			if tc.HasErr && err == nil {
				t.Error("Expected err, but does not found.")
			}
			if got != tc.Exp {
				t.Errorf("Expected: '%v', got: '%v' for: '%s'", tc.Exp, got, tc.Val)
			}
		})
	}
}

func TestIsUniqueCharactersMemorize(t *testing.T) {
	for _, tc := range getTestDataIsUniqueCharacters() {
		t.Run(tc.Name, func(t *testing.T) {
			got, err := IsUniqueCharactersMemorize(tc.Val)
			if tc.HasErr && err == nil {
				t.Error("Expected err, but does not found.")
			}
			if got != tc.Exp {
				t.Errorf("Expected: '%v', got: '%v' for: '%s'", tc.Exp, got, tc.Val)
			}
		})
	}
}

func BenchmarkIsUniqueCharactersIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range getTestDataIsUniqueCharacters() {
			res, _ := IsUniqueCharactersIterative(tc.Val)
			resultIsUnique = res
		}
	}
}

func BenchmarkIsUniqueCharactersMemorize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range getTestDataIsUniqueCharacters() {
			res, _ := IsUniqueCharactersMemorize(tc.Val)
			resultIsUnique = res
		}
	}
}

func getTestDataIsUniqueCharacters() []testDataIsUniqueCharacters {
	return []testDataIsUniqueCharacters{
		{
			"IllegalEncodingUTF8-unique",
			"\x80日本語漢字かんじ",
			false,
			true,
		},
		{
			"EncodingUTF8-NOT-unique",
			"日本語語漢字かんじ",
			false,
			false,
		},
		{
			"EncodingUTF8-unique",
			"日本語漢字かんじ",
			true,
			false,
		},
		{
			"ShortString-unique",
			"abdefghijk",
			true,
			false,
		},
		{
			"ShortString-NOT-unique",
			"abdefghijak",
			false,
			false,
		},
		{
			"TestEmptyString-unique",
			"",
			true,
			false,
		},
	}
}
