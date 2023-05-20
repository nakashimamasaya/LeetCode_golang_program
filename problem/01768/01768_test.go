package _01768

import "testing"

func TestSample(t *testing.T) {
	word1 := []string{
		"abc",
		"ab",
		"abcd",
	}

	word2 := []string{
		"pqr",
		"pqrs",
		"pq",
	}

	ans := []string{
		"apbqcr",
		"apbqrs",
		"apbqcd",
	}

	for i, a := range ans {
		if result := mergeAlternately(word1[i], word2[i]); result != a {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
