package _00205

import "testing"

func TestSample(t *testing.T) {
	ss := []string{
		"egg",
		"foo",
		"paper",
	}

	ts := []string{
		"add",
		"bar",
		"title",
	}

	ans := []bool{
		true,
		false,
		true,
	}

	for i := 0; i < len(ss); i++ {
		if result := isIsomorphic(ss[i], ts[i]); result != ans[i] {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, ans[i])
		}
	}
}