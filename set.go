package goset

type Set struct {
	set map[interface{}]bool
}

func NewSet() *Set {
	return &Set{make(map[interface{}]bool)}
}

func (set *Set) Add(items ...interface{}) bool {
	for _, item := range items {
		set.set[item] = true
	}
	return true
}

func (set *Set) Size() int {
	return len(set.set)
}

func (set *Set) IsMember(i interface{}) bool {
	_, exists := set.set[i]
	return exists
}

func (set *Set) Del(i interface{}) {
	delete(set.set, i)
}

func (A *Set) Union(B *Set) *Set {
	newSet := NewSet()
	for key, _ := range A.set {
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
		if A.IsMember(key) {
			newSet.Add(key)
		}
	}
	return newSet
}

func (A *Set) Intersect(B *Set) *Set {
	ACard := A.Size()
	BCard := B.Size()
	if ACard > BCard {
		return A.merge(B)
	} else {
		return B.merge(A)
	}
}
