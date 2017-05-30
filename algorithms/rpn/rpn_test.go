package rpn

import "testing"

func TestEvalRPN(t *testing.T) {
	scenarios := []struct {
		tokens []string
		result int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
	}
	for _, scenario := range scenarios {
		actual := evalRPN(scenario.tokens)
		if actual != scenario.result {
			t.Fatalf("tokens: %v, want: %v, actual: %v", scenario.tokens, scenario.result, actual)
		}
	}
}
