package _0012

// 12. Integer to Roman
// https://leetcode.com/problems/integer-to-roman/

func intToRoman(num int) string {
	ans := ""
	dic := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	nums := []int{
		1000, 100, 10, 1,
	}

	for _, n := range nums {
		tmp := num / n
		num = num % n
		if tmp == 9 || tmp == 4 {
			ans += dic[n] + dic[n*(tmp+1)]
			continue
		}
		if tmp >= 5 {
			ans += dic[5*n]
			tmp -= 5
		}

		for i := 0; i < tmp; i++ {
			ans += dic[n]
		}
	}
	return ans
}
