package _0100

import "testing"

func TestSample1(t *testing.T) {
	pL := TreeNode{Val: 2}
	pR := TreeNode{Val: 3}
	p := TreeNode{Val: 1, Left: &pL, Right: &pR}

	qL := TreeNode{Val: 2}
	qR := TreeNode{Val: 3}
	q := TreeNode{Val: 1, Left: &qL, Right: &qR}

	result := IsSameTree(&p, &q)
	ans := true

	if result != ans {
		t.Error("miss sample1")
	}
}

func TestSample2(t *testing.T) {
	pL := TreeNode{Val: 2}
	p := TreeNode{Val: 1, Left: &pL}

	qR := TreeNode{Val: 2}
	q := TreeNode{Val: 1, Right: &qR}

	result := IsSameTree(&p, &q)
	ans := false
	if result != ans {
		t.Error("miss sample2")
	}
}

func TestSample3(t *testing.T) {
	pL := TreeNode{Val: 2}
	pR := TreeNode{Val: 1}
	p := TreeNode{Val: 1, Left: &pL, Right: &pR}

	qL := TreeNode{Val: 1}
	qR := TreeNode{Val: 2}
	q := TreeNode{Val: 1, Left: &qL, Right: &qR}

	result := IsSameTree(&p, &q)
	ans := false

	if result != ans {
		t.Error("miss sample3")
	}
}
