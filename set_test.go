package goset_test

import (
	"testing"

	"github.com/abishekk92/goset"
)

func TestAddToSet(t *testing.T) {
	setStr := goset.NewSet()
	result := setStr.Add("foobar")
	if result != true {
		t.Error("Expected", true, "but got", result)
	}
}

func TestMultiAdd(t *testing.T) {
	setStr := goset.NewSet()
	result := setStr.Add("foobar", "zigzag", "zipzap")
	if result != true {
		t.Error("Expected", true, "but got", result)
	}

}

func TestSizeofSet(t *testing.T) {
	setStr := goset.NewSet()
	setStr.Add("foobar")
	result := setStr.Size()
	if result != 1 {
		t.Error("Expected", 1, "but got", result)
	}
}

func TestIsMember(t *testing.T) {
	setStr := goset.NewSet()
	setStr.Add("foobar")
	result := setStr.IsMember("foobar")
	if result != true {
		t.Error("Expected", true, "but got", result)
	}
}

func TestDel(t *testing.T) {
	setStr := goset.NewSet()
	setStr.Add("foo")
	setStr.Add("bar")
	setStr.Del("foo")
	result := setStr.IsMember("foo")
	if result != false {
		t.Error("Expected", false, "but got", result)
	}
	setSize := setStr.Size()
	if setSize != 1 {
		t.Error("Expected", 1, "but got", setSize)
	}
}

func TestSetUnion(t *testing.T) {
	A := goset.NewSet()
	B := goset.NewSet()
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
	A := goset.NewSet()
	B := goset.NewSet()
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
