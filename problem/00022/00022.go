package _0022

// 22. Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/

func GenerateParenthesis(n int) []string {
	return memoGenerateParenthesis(n, map[int][]string{0: {}, 1: {"()"}})
}

func memoGenerateParenthesis(n int, memo map[int][]string) []string {
	if v, ok := memo[n]; ok {
		return v
	}

	resultMap := map[string]int{}

	for i := 1; i <= n/2; i++ {
		for _, v1 := range memoGenerateParenthesis(n-i, memo) {
			for _, v2 := range memoGenerateParenthesis(i, memo) {
				resultMap[v1+v2] = 0
				resultMap[v2+v1] = 0
			}
			if i == 1 {
				resultMap["("+v1+")"] = 0
			}
		}
	}

	var result []string
	for k := range resultMap {
		result = append(result, k)
	}

	memo[n] = result
	return result
}
