package _00242

import "testing"

func TestSample(t *testing.T) {
	ss := []string{
		"anagram",
		"rat",
	}

	ts := []string{
		"nagaram",
		"car",
	}

	ans := []bool{
		true,
		false,
	}

	for i := 0; i < len(ss); i++ {
		result := isAnagram(ss[i], ts[i])
		if result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
