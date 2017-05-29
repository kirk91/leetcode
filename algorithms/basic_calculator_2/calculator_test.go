package calculator

import "testing"

func TestCalculate(t *testing.T) {
	cases := []struct {
		s string
		n int
	}{
		{"3+2*2", 7},
		{" 3/2", 1},
		{" 3+5 / 2", 5},
	}
	for i := range cases {
		actual := calculate(cases[i].s)
		if cases[i].n != actual {
			t.Fatalf("statement %s want: %d, but got: %d", cases[i].s, cases[i].n, actual)
		}
	}
}
