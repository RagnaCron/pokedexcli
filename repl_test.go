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
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "UPPERCASE lowercase CamelCase",
			expected: []string{"uppercase", "lowercase", "camelcase"},
		},
		{
			input:    "Some more Test Cases",
			expected: []string{"some", "more", "test", "cases"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual length: '%v' does not match expected length: '%v'", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual word: '%v' does not match expected word: '%v'", word, expectedWord)
			}
		}
	}
}
