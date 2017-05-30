package number

func singleNumber(nums []int) []int {
	diff := 0
	for _, num := range nums {
		diff ^= num
	}
	diff &= -diff // get last 1 bit
	ret := make([]int, 2)
	for _, num := range nums {
		if diff&num == 0 {
			ret[0] ^= num
		} else {
			ret[1] ^= num
		}
	}
	return ret
}
