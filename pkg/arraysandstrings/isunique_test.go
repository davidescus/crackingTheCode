package arraysandstrings

import "testing"

var result bool

type testData struct {
	Name   string
	Val    string
	Exp    bool
	HasErr bool
}

func TestIsUniqueCharactersIterative(t *testing.T) {
	for _, tc := range getTestData() {
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
	for _, tc := range getTestData() {
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
		for _, tc := range getTestData() {
			res, _ := IsUniqueCharactersIterative(tc.Val)
			result = res
		}
	}
}

func BenchmarkIsUniqueCharactersMemorize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range getTestData() {
			res, _ := IsUniqueCharactersMemorize(tc.Val)
			result = res
		}
	}
}

func getTestData() []testData {
	return []testData{
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
