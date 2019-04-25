package main

import "testing"

func TestCalculatePrimacy(t *testing.T) {
	cases := []struct {
		input  int
		wanted bool
	}{
		{1007, false},
		{10357, true},
	}
	for _, c := range cases {
		result := CalculatePrimacy(c.input)
		if result != c.wanted {
			t.Errorf("CalculatePrimacy(%v) == %v, wanted %v", c.input, result, c.wanted)
		}
	}

}
