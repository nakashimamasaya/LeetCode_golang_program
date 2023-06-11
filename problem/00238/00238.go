package _00238

// 238. Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self

func productExceptSelf(nums []int) []int {
	zeroF, product := 0, 1
	ans := make([]int, len(nums))
	for _, n := range nums {
		if n == 0 {
			zeroF++
			continue
		}
		product *= n
	}

	for i := 0; i < len(nums) && zeroF < 2; i++ {
		if zeroF == 0 {
			ans[i] = product / nums[i]
		} else if zeroF == 1 && nums[i] == 0 {
			ans[i] = product
			break
		}
	}
	return ans
}
