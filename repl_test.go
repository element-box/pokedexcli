package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " PIKACHU Bulbasaur squirtle",
			expected: []string{"pikachu", "bulbasaur", "squirtle"},
		},
		{
			input:    "Charmander is best",
			expected: []string{"charmander", "is", "best"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The actual length of slices %d doesn't match the expected length %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word %s does not match word %s", word, expectedWord)
			}
		}
	}
}
