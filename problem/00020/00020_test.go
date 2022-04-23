package _0020

import "testing"

func TestSample1(t *testing.T) {
	s := "()"
	if IsValid((s)) != true {
		t.Error("sample1 miss")
	}
}

func TestSample2(t *testing.T) {
	s := "()[]{}"
	if IsValid(s) != true {
		t.Error("sample2 miss")
	}
}

func TestSample3(t *testing.T) {
	s := "(]"
	if IsValid(s) != false {
		t.Error("sample3 miss")
	}
}
