package _02908

// 2908. Minimum Sum of Mountain Triplets I
// https://leetcode.com/problems/minimum-sum-of-mountain-triplets-i/description/

func minimumSum(nums []int) int {
	ans := -1
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] < nums[j] && nums[k] < nums[j] &&
					(ans == -1 || ans > nums[i]+nums[j]+nums[k]) {
					ans = nums[i] + nums[j] + nums[k]
				}
			}
		}

	}
	return ans
}
