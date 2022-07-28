package fundamentalstest

import "testing"

const spanish = "Spanish"
const french = "French"
const englishPrefixHallo = "Hallo, "
const spanishPrefixHallo = "Hola, "
const frenchPrefixHallo = "Bonjour, "

func gettingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefixHallo
	case french:
		prefix = frenchPrefixHallo
	default:
		prefix = englishPrefixHallo
	}

	return
}

func SayHallo(name string, language string) string {
	prefix := gettingPrefix(language)

	if name == "" {
		name = "World"
	}

	return prefix + name
}

func TestHallo(t *testing.T) {
	testCases := []struct {
		desc     string
		name     string
		language string
		expected string
	}{
		{
			desc:     "empty params",
			name:     "",
			expected: "Hallo, World",
		},
		{
			desc:     "params exist",
			name:     "Dhany",
			expected: "Hallo, Dhany",
		},
		{
			desc:     "languange spanish",
			name:     "Dhany",
			language: "Spanish",
			expected: "Hola, Dhany",
		},
		{
			desc:     "languange french",
			name:     "Dhany",
			language: "French",
			expected: "Bonjour, Dhany",
		},
		{
			desc:     "others languange",
			name:     "Dhany",
			language: "Italy",
			expected: "Hallo, Dhany",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := SayHallo(tC.name, tC.language)
			if tC.expected != result {
				t.Error("expected:", tC.expected, "got:", result)
			}
		})
	}
}
