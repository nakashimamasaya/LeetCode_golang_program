package _02810

// 2810. Faulty Keyboard
// https://leetcode.com/problems/faulty-keyboard/

func finalString(s string) string {
	ans := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			for j, k := 0, len(ans)-1; j < k; j, k = j+1, k-1 {
				ans[j], ans[k] = ans[k], ans[j]
			}
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
