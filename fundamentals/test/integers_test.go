package fundamentalstest

import "testing"

func Add(x, y int) int {
	return x + y
}

func TestAdder(t *testing.T) {
	testCases := []struct {
		desc     string
		value1   int
		value2   int
		expected int
	}{
		{
			desc:     "1 + 1",
			value1:   1,
			value2:   1,
			expected: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := Add(tC.value1, tC.value2)
			if tC.expected != result {
				t.Error("expected:", tC.expected, "got:", result)
			}
		})
	}
}
