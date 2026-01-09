package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
    // anonymous struct
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hi my name is Khang",
			expected: []string{"hi", "my", "name", "is", "Khang"},
		},

        {
            input: "   git     commit ",
            expected: []string{"git", "commit"},
        },

        {
            input: "hey",
            expected: []string{"hey"},
        },
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
            if word != expectedWord {
                fmt.Errorf("wrong ass impl")
            }
	 	}
	}
}
