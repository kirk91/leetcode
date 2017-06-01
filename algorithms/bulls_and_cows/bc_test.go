package bc

import "testing"

func TestGetHint(t *testing.T) {
	scenarios := []struct {
		secret string
		guess  string
		result string
	}{
		{"1807", "7810", "1A3B"},
		{"1123", "0111", "1A1B"},
	}
	for _, scenario := range scenarios {
		actual := getHint(scenario.secret, scenario.guess)
		if actual != scenario.result {
			t.Fatalf("secret: %s guess: %s, want: %v got: %v", scenario.secret, scenario.guess, scenario.result, actual)
		}
	}
}
