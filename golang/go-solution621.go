package main


// 任务调度器
func leastInterval(tasks []byte, n int) int {
	charSlice := [26]int{}
	m := 0
	cnt := 0
	for _, v := range tasks {
		charSlice[v - 'A']++
		if charSlice[v - 'A'] > m {
			m = charSlice[v - 'A']
			cnt = 1
		} else if charSlice[v - 'A'] == m {
			cnt++
		}
	}

	if n == 0 || (m - 1) * (n + 1) + cnt < len(tasks) {
		return len(tasks)
	}

	return (m - 1) * (n + 1) + cnt
}

// ["A","A","A","B","B","B"]
