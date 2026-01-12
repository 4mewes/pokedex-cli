package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
		}{
			{
				input:    "  hello  world  ",
				expected: []string{"hello", "world"},
			},
			{
				input: "Charmander Bulbasaur",
				expected: []string{"Charmander", "Bulbasaur"},
			},
			{
				input: "Ditto",
				expected: []string{"Ditto"},
			},
	}

	for _, c := range cases {
			actual := cleanInput(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("actual len does not match expected len")		
			}
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord {
					t.Errorf("words do not match")
				}
	}
}
}
