package fundamentalstest

import (
	"strings"
	"testing"
)

func Repeat(character string, repeat int) string {
	result := ""
	for i := 0; i < repeat; i++ {
		result += character
	}
	return result
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		desc      string
		character string
		repeat    int
		expected  string
	}{
		{
			desc:      "repeat a",
			character: "a",
			repeat:    5,
			expected:  strings.Repeat("a", 5),
		},
		{
			desc:      "repeat ",
			character: "",
			repeat:    5,
			expected:  strings.Repeat("", 5),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := Repeat(tC.character, tC.repeat)
			if result != tC.expected {
				t.Error("expected:", tC.expected, "got:", result)
			}
		})
	}
}
