README
======

This is an Naive Implementation of a map data structure with no  dynamic sizing in golang. The map is 
implemented in naive_map.go and tested in naive_map_test.go. In order to build and test:

  $ go build
  $ go test -v

If intended to use in external functionality:

```go

import "NaiveMap"

```

Should be designed to deal with concurrency, but might behave with issues with serious parrallelization!

to create new NaiveMap, use the constructor as following:

```go

map := NewNaiveMap(size)

```

Where size is the fixed size of the map.

