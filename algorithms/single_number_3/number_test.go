package number

import (
	"reflect"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{[]int{2, 1, 2, 3, 4, 1}, []int{4, 3}},
		{[]int{1, 2, 1, 3, 2, 5}, []int{5, 3}},
	}
	for i := range cases {
		want := cases[i].output
		actual := singleNumber(cases[i].input)
		if !reflect.DeepEqual(actual, want) {
			t.Fatalf("Input: %v, want: %v, actual: %v", cases[i].input, want, actual)
		}
	}
}
