goset
=====

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
        newSet.Add("a") //Adds a new element to the set
        newSet.Add("b")
        newSet.del("b") //Rmeoves an element from the set
        newSet.isMember("a") //Checks if 'a' is a member of the set
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
        A.add("a")
        A.add("b")
        B.add("a")
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
        A.add("a")
        A.add("b")
        B.add("a")
        intersectionSet := A.Intersect(B)
    }
```
