package arraysandstrings

import (
	"bytes"
	"testing"
)

type testDataURLify struct {
	name string
	str  []byte
	exp  []byte
	len  int
}

func TestURLify(t *testing.T) {
	for _, tc := range getTestDataURLify() {
		t.Run(tc.name, func(t *testing.T) {
			got := URLifyAscii(tc.str, tc.len)
			if !bytes.Equal(got, tc.exp) {
				t.Errorf("%s: Exp: %s, Got: %s", tc.name, tc.exp, got)
			}
		})
	}
}

func getTestDataURLify() []testDataURLify {
	return []testDataURLify{
		{
			name: "Empty",
			str:  []byte(""),
			exp:  []byte(""),
			len:  0,
		},
		{
			name: "NoSpaces",
			str:  []byte("John"),
			exp:  []byte("John"),
			len:  4,
		},
		{
			name: "OneSpace",
			str:  []byte("John Doe  "),
			exp:  []byte("John%20Doe"),
			len:  8,
		},
		{
			name: "MultipleSpaces",
			str:  []byte("John Doe Pamela    "),
			exp:  []byte("John%20Doe%20Pamela"),
			len:  15,
		},
	}
}
