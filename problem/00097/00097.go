package _00097

import "reflect"

// 97. Interleaving String
// https://leetcode.com/problems/interleaving-string/

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) == 0 && len(s2) == 0 && len(s3) == 0 {
		return true
	} else if len(s1)+len(s2) != len(s3) {
		return false
	} else if !isMatchCountCaracter(s1+s2, s3) {
		return false
	}

	return isA(s1, s2, s3)
}

func isA(s1, s2, s3 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return s1 == s3 || s2 == s3
	}

	ans := false
	if s1[0] == s3[0] {
		ans = isA(s1[1:], s2, s3[1:])
	}
	if !ans && s2[0] == s3[0] {
		ans = isA(s1, s2[1:], s3[1:])
	}
	return ans
}

func isMatchCountCaracter(s1, s2 string) bool {
	m, m2 := map[string]int{}, map[string]int{}
	for i := 0; i < len(s1); i++ {
		m[string(s1[i])]++
	}
	for i := 0; i < len(s2); i++ {
		m2[string(s2[i])]++
	}

	return reflect.DeepEqual(m, m2)
}
