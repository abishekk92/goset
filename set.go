package goset

import (
	"sync"
)

type Set struct {
	set  map[interface{}]bool
	sync bool
	sync.RWMutex
}

func NewSet(sync bool) *Set {
	return &Set{set: make(map[interface{}]bool), sync: sync}
}

func (set *Set) Add(items ...interface{}) {
	if set.sync {
		set.Lock()
		defer set.Unlock()
	}
	for _, item := range items {
		set.set[item] = true
	}
}

func (set *Set) Size() int {
	if set.sync {
		set.Lock()
		defer set.Unlock()
	}
	return len(set.set)
}

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

func (set *Set) Remove(i interface{}) {
	if set.sync {
		set.Lock()
		delete(set.set, i)
		set.Unlock()
	} else {
		delete(set.set, i)
	}
}

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

func (A *Set) merge(B *Set) *Set {
	newSet := NewSet(A.sync)
	for key := range B.set {
		if A.IsMember(key) {
			newSet.Add(key)
		}
	}
	return newSet
}

func (A *Set) Intersect(B *Set) *Set {
	if A.Size() > B.Size() {
		return A.merge(B)
	} else {
		return B.merge(A)
	}
}

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

func (A *Set) Difference(B *Set) *Set {
	newSet := NewSet(A.sync)
	for key := range A.set {
		if B.IsMember(key) != true {
			newSet.Add(key)
		}
	}
	return newSet
}
