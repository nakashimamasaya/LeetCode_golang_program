package _00125

import "testing"

func TestSample(t *testing.T) {
	s := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		" ",
	}

	ans := []bool{
		true,
		false,
		true,
	}

	for i, a := range ans {
		if result := isPalindrome(s[i]); a != result {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
