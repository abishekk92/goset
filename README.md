goset
=====

[![Build Status](https://travis-ci.org/abishekk92/goset.svg?branch=master)](https://travis-ci.org/abishekk92/goset) [![codecov](https://codecov.io/gh/abishekk92/goset/branch/master/graph/badge.svg)](https://codecov.io/gh/abishekk92/goset) [![GoDoc](https://godoc.org/github.com/abishekk92/goset?status.svg)](https://godoc.org/github.com/abishekk92/goset) [![Go Report Card](https://goreportcard.com/badge/github.com/abishekk92/goset)](https://goreportcard.com/report/github.com/abishekk92/goset) [![experimental](http://badges.github.io/stability-badges/dist/experimental.svg)](http://github.com/badges/stability-badges)

Set implementation for golang

_To install goset_

```bash
    go get github.com/abishekk92/goset
```

_To Create a new set_

```go
    import (
        "github.com/abishekk92/goset"
        )

    func main() {
        newSet := goset.NewSet() //Creates a new set
        newSet.Add("a", "b") //Adds a new element to the set
        newSet.Remove("b") //Removes an element from the set
        newSet.IsMember("a") //Checks if 'a' is a member of the set
    }
```


_To Union two sets_

```go
    import (
        "github.com/abishekk92/goset"
        )

    func main() {
        A := goset.NewSet()
        B := goset.NewSet()
        A.Add("a", "b")
        B.Add("a")
        uninonSet := A.Union(B)
    }
```

_To Intersect two sets_

```go
    import (
        "github.com/abishekk92/goset"
        )

    func main() {
        A := goset.NewSet()
        B := goset.NewSet()
        A.Add("a", "b")
        B.Add("a")
        intersectionSet := A.Intersect(B)
    }
```

_To get an array of elements_

```go
    import (
        "github.com/abishekk92/goset"
        )

    func main() {
        A := goset.NewSet()
        A.Add("a", "b")
        items := A.ToArray()
    }
```

_To find set difference_

```go
    import (
        "github.com/abishekk92/goset"
        )

    func main() {
        A := goset.NewSet()
        A.Add("a", "b", "c", "d")
        B := goset.NewSet()
        B.Add("a", "b")
        C := A.Difference(B) // result would be C{"c", "d"}
    }
```
