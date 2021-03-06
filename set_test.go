package goset_test

import (
	"testing"

	"github.com/abishekk92/goset"
)

func TestAddToSet(t *testing.T) {
	setStr := goset.NewSet(false)
	setStr.Add("foobar")
	result := setStr.IsMember("foobar")
	if result != true {
		t.Error("Expected", true, "but got", result)
	}
}

func TestMultiAdd(t *testing.T) {
	setStr := goset.NewSet(false)
	items := []interface{}{"foobar", "zigzag", "zipzap"}
	setStr.Add(items...)
	for _, item := range items {
		result := setStr.IsMember(item)
		if result != true {
			t.Error("Expected", true, "but got", result)
		}
	}
}

func TestSizeofSet(t *testing.T) {
	setStr := goset.NewSet(false)
	setStr.Add("foobar")
	result := setStr.Size()
	if result != 1 {
		t.Error("Expected", 1, "but got", result)
	}
}

func TestIsMember(t *testing.T) {
	setStr := goset.NewSet(false)
	setStr.Add("foobar")
	result := setStr.IsMember("foobar")
	if result != true {
		t.Error("Expected", true, "but got", result)
	}
}

func TestRemove(t *testing.T) {
	setStr := goset.NewSet(false)
	setStr.Add("foo", "bar")
	setStr.Remove("foo")
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
	A := goset.NewSet(false)
	B := goset.NewSet(false)
	A.Add("a", "b", "c", "d")
	B.Add("a", "b")
	ACard := A.Size()
	BCard := B.Size()
	result := A.Union(B)
	resultSize := result.Size()
	if resultSize < BCard || resultSize > ACard+BCard {
		t.Error("The Union of Set A and B has been compromised")
	}
}

func TestSetIntersection(t *testing.T) {
	A := goset.NewSet(false)
	B := goset.NewSet(false)
	A.Add("a", "b", "c", "d")
	B.Add("a", "b")
	ACard := A.Size()
	BCard := B.Size()
	result := A.Intersect(B)
	resultSize := result.Size()
	if resultSize < BCard || resultSize > ACard-BCard {
		t.Error("The Intersection of Set A and B has been compromised")
	}
}

func TestSetIntersection2(t *testing.T) {
	A := goset.NewSet(false)
	B := goset.NewSet(false)
	A.Add("a", "b")
	B.Add("a", "b", "c", "d")
	ACard := A.Size()
	BCard := B.Size()
	result := A.Intersect(B)
	resultSize := result.Size()
	if resultSize < ACard || resultSize > BCard-ACard {
		t.Error("The Intersection of Set A and B has been compromised")
	}
}

func TestSetToArray(t *testing.T) {
	A := goset.NewSet(false)
	items := []interface{}{"foobar", "zigzag", "zipzap"}
	A.Add(items...)
	keysArray := A.ToArray()
	if len(keysArray) != A.Size() {
		t.Error("Information seems to have been lost while converting to array")
	}
	if len(keysArray) != len(items) {
		t.Error("Information seems to have been lost while converting to array")
	}
	for i := 0; i < len(keysArray); i++ {
		if keysArray[i] == nil {
			t.Error("Information seems to have been lost while converting to array")
		}
	}
}

func TestSetDifference(t *testing.T) {
	A := goset.NewSet(false)
	B := goset.NewSet(false)
	A.Add("a", "b", "c", "d")
	B.Add("a", "b")
	C := A.Difference(B)
	if C.Size() != 2 {
		t.Error("Expected difference to have 2 elements")
	}
	if !(C.IsMember("c") && C.IsMember("d")) {
		t.Error("Expected the result to contain elements 'c' and 'd'")
	}
}

func TestSyncAddToSet(t *testing.T) {
	setStr := goset.NewSet(true)
	setStr.Add("foobar")
	result := setStr.IsMember("foobar")
	if result != true {
		t.Error("Expected", true, "but got", result)
	}
}

func TestSyncMultiAdd(t *testing.T) {
	setStr := goset.NewSet(true)
	items := []interface{}{"foobar", "zigzag", "zipzap"}
	setStr.Add(items...)
	for _, item := range items {
		result := setStr.IsMember(item)
		if result != true {
			t.Error("Expected", true, "but got", result)
		}
	}
}

func TestSyncSizeofSet(t *testing.T) {
	setStr := goset.NewSet(true)
	setStr.Add("foobar")
	result := setStr.Size()
	if result != 1 {
		t.Error("Expected", 1, "but got", result)
	}
}

func TestSyncIsMember(t *testing.T) {
	setStr := goset.NewSet(true)
	setStr.Add("foobar")
	result := setStr.IsMember("foobar")
	if result != true {
		t.Error("Expected", true, "but got", result)
	}
}

func TestSyncRemove(t *testing.T) {
	setStr := goset.NewSet(true)
	setStr.Add("foo", "bar")
	setStr.Remove("foo")
	result := setStr.IsMember("foo")
	if result != false {
		t.Error("Expected", false, "but got", result)
	}
	setSize := setStr.Size()
	if setSize != 1 {
		t.Error("Expected", 1, "but got", setSize)
	}
}

func TestSyncSetUnion(t *testing.T) {
	A := goset.NewSet(true)
	B := goset.NewSet(true)
	A.Add("a", "b", "c", "d")
	B.Add("a", "b")
	ACard := A.Size()
	BCard := B.Size()
	result := A.Union(B)
	resultSize := result.Size()
	if resultSize < BCard || resultSize > ACard+BCard {
		t.Error("The Union of Set A and B has been compromised")
	}
}

func TestSyncSetIntersection(t *testing.T) {
	A := goset.NewSet(true)
	B := goset.NewSet(true)
	A.Add("a", "b", "c", "d")
	B.Add("a", "b")
	ACard := A.Size()
	BCard := B.Size()
	result := A.Intersect(B)
	resultSize := result.Size()
	if resultSize < BCard || resultSize > ACard-BCard {
		t.Error("The Union of Set A and B has been compromised")
	}
}

func TestSyncSetIntersection2(t *testing.T) {
	A := goset.NewSet(true)
	B := goset.NewSet(true)
	A.Add("a", "b")
	B.Add("a", "b", "c", "d")
	ACard := A.Size()
	BCard := B.Size()
	result := A.Intersect(B)
	resultSize := result.Size()
	if resultSize < ACard || resultSize > BCard-ACard {
		t.Error("The Intersection of Set A and B has been compromised")
	}
}

func TestSyncSetToArray(t *testing.T) {
	A := goset.NewSet(true)
	items := []interface{}{"foobar", "zigzag", "zipzap"}
	A.Add(items...)
	keysArray := A.ToArray()
	if len(keysArray) != A.Size() {
		t.Error("Information seems to have been lost while converting to array")
	}
	if len(keysArray) != len(items) {
		t.Error("Information seems to have been lost while converting to array")
	}
	for i := 0; i < len(keysArray); i++ {
		if keysArray[i] == nil {
			t.Error("Information seems to have been lost while converting to array")
		}
	}
}

func TestSyncSetDifference(t *testing.T) {
	A := goset.NewSet(true)
	B := goset.NewSet(true)
	A.Add("a", "b", "c", "d")
	B.Add("a", "b")
	C := A.Difference(B)
	if C.Size() != 2 {
		t.Error("Expected difference to have 2 elements")
	}
	if !(C.IsMember("c") && C.IsMember("d")) {
		t.Error("Expected the result to contain elements 'c' and 'd'")
	}
}
