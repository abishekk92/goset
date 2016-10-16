package goset

import "sync"

type Set struct {
	sync.RWMutex
	set map[interface{}]bool
}

func NewSet() *Set {
	return &Set{set: make(map[interface{}]bool)}
}

func (set *Set) Add(i interface{}) bool {
	set.Lock()
	_, exists := set.set[i]
	set.set[i] = true
	set.Unlock()
	return !exists
}

func (set *Set) Size() int {
	set.RLock()
	defer set.RUnlock()

	return len(set.set)
}

func (set *Set) isMember(i interface{}) bool {
	set.RLock()
	_, exists := set.set[i]
	set.RUnlock()
	return exists
}

func (set *Set) del(i interface{}) {
	set.Lock()
	delete(set.set, i)
	set.Unlock()
}

func (A *Set) Union(B *Set) *Set {
	A.RLock()
	defer A.RUnlock()
	B.RLock()
	defer B.RUnlock()

	newSet := NewSet()
	for key := range A.set {
		newSet.Add(key)
	}

	for key, _ := range B.set {
		newSet.Add(key)
	}

	return newSet
}

func (A *Set) merge(B *Set) *Set {
	newSet := NewSet()
	for key, _ := range B.set {
		if A.isMember(key) {
			newSet.Add(key)
		}
	}
	return newSet
}

func (A *Set) Intersect(B *Set) *Set {
	A.RLock()
	ACard := A.Size()
	A.RUnlock()

	B.RLock()
	BCard := B.Size()
	B.RUnlock()

	if ACard > BCard {
		return A.merge(B)
	}

	return B.merge(A)
}
