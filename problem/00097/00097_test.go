package _00097

import "testing"

func TestSample(t *testing.T) {
	s1 := []string{
		"aabcc",
		"aabcc",
		"",
	}

	s2 := []string{
		"dbbca",
		"dbbca",
		"",
	}

	s3 := []string{
		"aadbbcbcac",
		"aadbbbaccc",
		"",
	}

	ans := []bool{
		true,
		false,
		true,
	}

	for i := 0; i < len(s1); i++ {
		if result := isInterleave(s1[i], s2[i], s3[i]); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}
