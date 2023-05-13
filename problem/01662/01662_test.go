package _01662

import "testing"

func TestSample(t *testing.T) {
	word1 := [][]string{
		{"ab", "c"},
		{"a", "cb"},
		{"abc", "d", "defg"},
	}

	word2 := [][]string{
		{"a", "bc"},
		{"ab", "c"},
		{"abcddefg"},
	}

	ans := []bool{
		true,
		false,
		true,
	}

	for i, a := range ans {
		if result := arrayStringsAreEqual(word1[i], word2[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i, result, a)
		}
	}
}
