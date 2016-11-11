package goset

import (
	"sync"
)

//Struct to hold set
type Set struct {
	set  map[interface{}]bool
	sync bool
	sync.RWMutex
}

//Construct NewSet.
//Set sync = true to use a threadsafe version.
func NewSet(sync bool) *Set {
	return &Set{set: make(map[interface{}]bool), sync: sync}
}

//Add elements to the set.
func (set *Set) Add(items ...interface{}) {
	if set.sync {
		set.Lock()
		defer set.Unlock()
	}
	for _, item := range items {
		set.set[item] = true
	}
}

//Number of unique elements in a set.
func (set *Set) Size() int {
	if set.sync {
		set.Lock()
		defer set.Unlock()
	}
	return len(set.set)
}

//Check if an element is a member of the set.
func (set *Set) IsMember(i interface{}) bool {
	if set.sync {
		set.Lock()
		_, exists := set.set[i]
		set.Unlock()
		return exists
	} else {
		_, exists := set.set[i]
		return exists
	}
}

//Remove an element from the set.
func (set *Set) Remove(i interface{}) {
	if set.sync {
		set.Lock()
		delete(set.set, i)
		set.Unlock()
	} else {
		delete(set.set, i)
	}
}

//Union two sets.
func (A *Set) Union(B *Set) *Set {
	newSet := NewSet(A.sync)
	for key := range A.set {
		newSet.Add(key)
	}
	for key := range B.set {
		newSet.Add(key)
	}

	return newSet
}

//Intersect two sets. (Private wrapper)
func (A *Set) intersect(B *Set) *Set {
	newSet := NewSet(A.sync)
	for key := range B.set {
		if A.IsMember(key) {
			newSet.Add(key)
		}
	}
	return newSet
}

//Intersect two sets.
func (A *Set) Intersect(B *Set) *Set {
	if A.Size() > B.Size() {
		return A.intersect(B)
	} else {
		return B.intersect(A)
	}
}

//Export elements of a set as an array.
func (A *Set) ToArray() []interface{} {
	keys := make([]interface{}, A.Size())
	index := 0
	if A.sync {
		A.Lock()
		defer A.Unlock()
	}
	for key := range A.set {
		keys[index] = key
		index++
	}
	return keys
}

//Set difference between two given sets. i.e Elements in A and are not in B.
func (A *Set) Difference(B *Set) *Set {
	newSet := NewSet(A.sync)
	for key := range A.set {
		if B.IsMember(key) != true {
			newSet.Add(key)
		}
	}
	return newSet
}
