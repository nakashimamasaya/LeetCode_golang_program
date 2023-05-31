package _02515

// 2515. Shortest Distance to Target String in a Circular Array
// https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/

func closetTarget(words []string, target string, startIndex int) int {
	ans, n := -1, len(words)
	for i, rightI, leftI := 0, startIndex, startIndex; i <= n/2; i, leftI, rightI = i+1, (leftI-1+n)%n, (rightI+1)%n {
		if words[rightI] == target ||
			words[leftI] == target {
			ans = i
			break
		}
	}

	return ans
}
