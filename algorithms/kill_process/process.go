package process

func killProcess(pid []int, ppid []int, kill int) []int {
	childs := make(map[int][]int)
	for i := range ppid {
		if _, ok := childs[ppid[i]]; !ok {
			childs[ppid[i]] = make([]int, 0, 10)
		}
		childs[ppid[i]] = append(childs[ppid[i]], pid[i])
	}

	var killed []int
	killAll(kill, childs, &killed)
	return killed
}

func killAll(pid int, childs map[int][]int, killed *[]int) {
	*killed = append(*killed, pid)
	for _, child := range childs[pid] {
		killAll(child, childs, killed)
	}
}
