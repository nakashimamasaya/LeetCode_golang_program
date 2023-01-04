package _00205

// 205. Isomorphic Strings
// https://leetcode.com/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	m, used := map[string]string{}, map[string]bool{}
	for i := 0; i < len(s); i++ {
		if m[string(s[i])] == "" && used[string(t[i])] == false {
			m[string(s[i])] = string(t[i])
			used[string(t[i])] = true
		} else if m[string(s[i])] != string(t[i]) {
			return false
		}
	}
	return true
}
