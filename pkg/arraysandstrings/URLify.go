package arraysandstrings

// URLifyAscii will receive a string and his real length
// (string is filled with empty spaces at the end) and will
// replace each space with "%20", works well only with ascii chars
// unicode chars can have more de 1 byte
// time complexity: O(n)
// space complexity: O(1)
func URLifyAscii(s []byte, trueLen int) []byte {

	var spaceOccurencies int
	for i := 0; i < trueLen; i++ {
		spaceOccurencies++
	}

	return []byte("")
}
