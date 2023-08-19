package _02739

// 2739. Total Distance Traveled
// https://leetcode.com/problems/total-distance-traveled/

func distanceTraveled(mainTank int, additionalTank int) int {
	ans := 0
	for mainTank >= 5 && additionalTank >= 1 {
		ans += 5
		mainTank, additionalTank = mainTank-4, additionalTank-1
	}
	return (ans + mainTank) * 10
}
