package main

import "fmt"

func main() {
	fmt.Println(findIntegers(5))
}

// func findIntegers(num int) int {
// var count int
// for i := 0; i <= num; i++ {
// if check(i) {
// count += 1
// }
// }
// return count
// }

// func check(n int) bool {
// i := 31
// for i > 0 {
// if n&(1<<uint(i)) != 0 && n&(1<<uint(i-1)) != 0 {
// return false
// }
// i = i - 1
// }
// return true
// }

func findIntegers(num int) int {
	f := make([]int, 32)
	f[0], f[1] = 1, 2
	for i := 2; i < 32; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	var sum, prev_bit int
	for i := 31; i >= 0; {
		if num&(1<<uint(i)) != 0 {
			sum += f[i]
			if prev_bit == 1 {
				sum--
				break
			}
			prev_bit = 1
		} else {
			prev_bit = 0
		}
		i -= 1
	}
	return sum + 1
}
