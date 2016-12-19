package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	m, n := len(haystack), len(needle)
	for i := 0; i <= m-n; i++ {
		j := 0
		for ; j < n; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == n {
			return i
		}
	}
	return -1
}

func main() {
}
