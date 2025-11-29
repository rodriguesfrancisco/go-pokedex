package main

import (
	"testing"
)

func TestClearInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "multiple spaces between words",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "single word with spaces",
			input:    "   pokemon   ",
			expected: []string{"pokemon"},
		},
		{
			name:     "no spaces",
			input:    "help",
			expected: []string{"help"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "only spaces",
			input:    "     ",
			expected: []string{},
		},
		{
			name:     "multiple words with tabs and newlines",
			input:    "hello\tworld\nfoo",
			expected: []string{"hello", "world", "foo"},
		},
		{
			name:     "command with arguments",
			input:    "catch pikachu",
			expected: []string{"catch", "pikachu"},
		},
		{
			name:     "mixed whitespace",
			input:    "  \t hello  \n  world \t ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "three words",
			input:    "one two three",
			expected: []string{"one", "two", "three"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)

			// Check length first
			if len(actual) != len(c.expected) {
				t.Errorf("expected %d words but got %d", len(c.expected), len(actual))
				return
			}

			// Check each word
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]

				if word != expectedWord {
					t.Errorf("expected %s but found %s", expectedWord, word)
				}
			}
		})
	}
}
