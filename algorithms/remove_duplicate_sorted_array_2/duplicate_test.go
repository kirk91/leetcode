package duplicate

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 3}
	ret := removeDuplicates(nums)
	if ret != 4 && reflect.DeepEqual(nums[:4], []int{1, 1, 2, 3}) {
		t.Fatal("failed")
	}
}
