package bc

import "fmt"

func getHint(secret string, guess string) string {
	var bulls, cows int
	numbers := make([]int, 10)
	n := len(secret)
	for i := 0; i < n; i++ {
		s := secret[i] - '0'
		g := guess[i] - '0'
		if s == g {
			bulls++
		} else {
			if numbers[s] < 0 {
				cows++
			}
			if numbers[g] > 0 {
				cows++
			}
			numbers[s]++
			numbers[g]--
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
