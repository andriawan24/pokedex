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
			input:    "    hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "you are sunshine !",
			expected: []string{"you", "are", "sunshine", "!"},
		},
	}

	for i, c := range cases {
		actual := cleanInput(c.input)

		if len(c.expected) != len(actual) {
			t.Errorf("Error on test case: %d, Length is not the same between expected: %d and actual: %d", i+1, len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Words don't match between %s and %s", word, expectedWord)
			}
		}
	}
}
