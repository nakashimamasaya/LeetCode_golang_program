package _00383

// 383. Ransom Note
// https://leetcode.com/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
	m1, m2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(ransomNote); i++ {
		m1[ransomNote[i]] += 1
	}
	for i := 0; i < len(magazine); i++ {
		m2[magazine[i]] += 1
	}
	for k, v := range m1 {
		if m2[k] < v {
			return false
		}
	}
	return true
}
