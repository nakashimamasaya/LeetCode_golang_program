package _0020

// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/
func IsValid(s string) bool {
	var closeList []string
	closeMap := map[string]string{"{": "}", "(": ")", "[": "]"}
	for i := 0; i < len(s); i++ {
		if v, ok := closeMap[string(s[i])]; ok {
			closeList = append(closeList, v)
		} else if len(closeList) == 0 || string(s[i]) != closeList[len(closeList)-1] {
			return false
		} else {
			closeList = closeList[:len(closeList)-1]
		}
	}
	return len(closeList) == 0
}
