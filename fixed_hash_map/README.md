README
======

This is an Linear Probing Implementation of a map data structure with no  dynamic sizing in golang. The map is 
implemented in fixed_hash_map.go and tested in fixed_hash_map_test.go. In order to build and test:

	$ go build
	$ go test -v

If intended to use in external functionality:

''''go
import "FixedHashMap"
''''

Should be designed to deal with concurrency, but might behave with issues with serious parrallelization!

to create new FixedHashMap, use the constructor as following:

''''go
map := NewFixedHashMap(size)
''''

Where size is the fixed size of the map.
