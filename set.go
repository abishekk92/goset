package goset

type Set struct {
	set map[interface{}]bool
}

func NewSet() *Set {
	return &Set{make(map[interface{}]bool)}
}

func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.set[item] = true
	}
}

func (set *Set) Size() int {
	return len(set.set)
}

func (set *Set) IsMember(i interface{}) bool {
	_, exists := set.set[i]
	return exists
}

func (set *Set) Remove(i interface{}) {
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
	if A.Size() > B.Size() {
		return A.merge(B)
	} else {
		return B.merge(A)
	}
}

func (A *Set) ToArray() []interface{} {
	keys := make([]interface{}, A.Size())
	index := 0
	for key, _ := range A.set {
		keys[index] = key
		index++
	}
	return keys
}

func (A *Set) Difference(B *Set) *Set {
	newSet := NewSet()
	for key, _ := range A.set {
		if B.IsMember(key) != true {
			newSet.Add(key)
		}
	}
	return newSet
}
