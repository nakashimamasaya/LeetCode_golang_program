package _00443

import "testing"

func TestSample(t *testing.T) {
	chars := [][]byte{
		[]byte("aabbccc"),
		[]byte("a"),
		[]byte("abbbbbbbbbbbb"),
	}

	ans := []int{
		6,
		1,
		4,
	}

	ansSlice := [][]byte{
		[]byte("a2b2c3"),
		[]byte("a"),
		[]byte("ab12"),
	}

	for i, a := range ans {
		if result := compress(chars[i]); a != result || string(ansSlice[i]) != string(chars[i][:result]) {
			t.Errorf("miss sample %d\n return %v\n answer %v\n chars %v\n ansChars %v\n", i+1, result, a, chars[i][:result], ansSlice[i])
		}
	}
}
