package _00661

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	img := [][][]int{
		{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}},
	}

	ans := [][][]int{
		{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
	}

	for i, a := range ans {
		if result := imageSmoother(img[i]); !reflect.DeepEqual(result, a) {
			t.Errorf("miss sample %d\n return %v\n answer %v", i+1, result, a)
		}
	}
}
