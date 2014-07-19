package goset

import (
	"testing"
)

func TestAddToSet(t *testing.T) {
	setStr := NewSet()
	result := setStr.Add("foobar")
	if result != true {
		t.Error("Expected", true, "but got", result)
	}
}

func TestSizeofSet(t *testing.T) {
	setStr := NewSet()
	setStr.Add("foobar")
	result := setStr.Size()
	if result != 1 {
		t.Error("Expected", 1, "but got", result)
	}
}

func TestIsMember(t *testing.T) {
	setStr := NewSet()
	setStr.Add("foobar")
	result := setStr.isMember("foobar")
	if result != true {
		t.Error("Expected", true, "but got", result)
	}
}


func TestSetUnion(t *testing.T) {
    A := NewSet()
    B := NewSet()
    A.Add("a")
    A.Add("b")
    A.Add("c")
    A.Add("d")
    B.Add("a")
    B.Add("b")
	ACard := A.Size()
	BCard := B.Size()
	result := A.Union(B)
	resultSize := result.Size()
	if resultSize < BCard || resultSize > ACard+BCard {
		t.Error("The Union of Set A and B has been compromised")
	}
}

func TestSetIntersection(t *testing.T) {
    A := NewSet()
    B := NewSet()
    A.Add("a")
    A.Add("b")
    A.Add("c")
    A.Add("d")
    B.Add("a")
    B.Add("b")
	ACard := A.Size()
	BCard := B.Size()
	result := A.Intersect(B)
	resultSize := result.Size()
	if resultSize < BCard || resultSize > ACard-BCard {
		t.Error("The Union of Set A and B has been compromised")
	}
}
