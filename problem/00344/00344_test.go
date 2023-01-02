package _00344

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	ss := [][]byte{
		{104, 101, 108, 108, 111},
		{72, 97, 110, 110, 97, 104},
	}

	ans := [][]byte{
		{111, 108, 108, 101, 104},
		{104, 97, 110, 110, 97, 72},
	}

	for i, _ := range ss {
		if reverseString(ss[i]); !reflect.DeepEqual(ss[i], ans[i]) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, ss[i], ans[i])
		}
	}
}
