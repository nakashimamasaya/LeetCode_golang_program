package _02432

// 2432. The Employee That Worked on the Longest Task
// https://leetcode.com/problems/the-employee-that-worked-on-the-longest-task/

func hardestWorker(n int, logs [][]int) int {
	ans, time, leaveTime := 0, 0, 0
	for _, log := range logs {
		if log[1]-time > leaveTime {
			ans, leaveTime = log[0], log[1]-time
		} else if log[1]-time == leaveTime && ans > log[0] {
			ans = log[0]
		}
		time = log[1]
	}
	return ans
}
