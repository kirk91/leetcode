package process

import (
	"reflect"
	"testing"
)

func TestKillProcess(t *testing.T) {
	cases := []struct {
		pid    []int
		ppid   []int
		kill   int
		killed []int
	}{
		{[]int{1, 3, 10, 5}, []int{3, 0, 5, 3}, 5, []int{5, 10}},
	}

	for i := range cases {
		killed := killProcess(cases[i].pid, cases[i].ppid, cases[i].kill)
		if !reflect.DeepEqual(killed, cases[i].killed) {
			t.Fatalf("pid: %v ppid: %v kill: %v, want: %v, actual: %v", cases[i].pid, cases[i].ppid, cases[i].killed, killed)
		}
	}
}
