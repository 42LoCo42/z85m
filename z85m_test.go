package z85m

import "testing"

var vectors = [][]string{
	{"\xCA", ":]..F3"},
	{"\xCA\xFE", "+kICz2"},
	{"\xCA\xFE\xBA", "+kO%M1"},
	{"\xCA\xFE\xBA\xBE", "+kO#^"},
	{"\xCA\xFE\xBA\xBE\xDE\xAD\xBE\xEF\xFE\xED\xD0\x0D\xBA\xAD", "+kO#^?MsJX@{w5wX#=2n2"},
	{"\xCA\xFE\xBA\xBE\xDE\xAD\xBE\xEF\xFE\xED\xD0\x0D\xBA\xAD\xF0\x0D", "+kO#^?MsJX@{w5wX#>Dh"},
}

func TestEncode(t *testing.T) {
	for _, vector := range vectors {
		enc, err := Encode([]byte(vector[0]))
		if err != nil {
			panic(err)
		}
		if vector[1] != string(enc) {
			t.Log("Expected", vector[1], "but got", string(enc))
			t.Fail()
		}
	}
}

func TestDecode(t *testing.T) {
	for _, vector := range vectors {
		dec, err := Decode([]byte(vector[1]))
		if err != nil {
			panic(err)
		}
		if vector[0] != string(dec) {
			t.Log("Expected", vector[0], "but got", string(dec))
			t.Fail()
		}
	}
}
