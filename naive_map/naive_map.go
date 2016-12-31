// Package NaiveMap is an implementation of a fixed hash-map under the specifications outlied @  kpcbfellows.com/engineering/apply.
package NaiveMap

import (
	"log"
	"sync"
)

// NaiveMap is the struct initialization of the fixed size hash map.
type NaiveMap struct {
	// For thread safety
	lock sync.Mutex
	// Array of string keys for the hash map
	keys []string
	// Array of object values assigned to keys by index for the map.
	values []interface{}
	// Maximium number of <key, value> pairs possible to fit in map.
	capacity int
	// Number of elements currently in map
	count int
}

// NewNaiveMap creates the FixedSizeHashMap using the input construction size.
func NewNaiveMap(size int) *NaiveMap {
	initial := &NaiveMap{
		sync.Mutex{},
		make([]string, size),
		make([]interface{}, size),
		size,
		0,
	}
	return initial

}

// Set stores the <Key, Value> pair in the NaiveMap and returns a boolean if the operation was successful.
func (f *NaiveMap) Set(key string, value interface{}) bool {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.capacity <= f.count {
		log.Print("At maximum capacity!")
		return false
	}

	if key == "" {
		log.Print("Invalid input.")
		return false
	}
	// Check if they key has been set, if so, put value with key.
	for i, _ := range f.keys {
		if f.keys[i] == key {
			f.values[i] = value
			return true
		}
	}

	// Make sure there's an an empty spot, then add.
	for i, _ := range f.keys {
		if f.keys[i] == "" {
			f.keys[i] = key
			f.values[i] = value
			f.count++
			return true
		}
	}
	return false
}

// Get returns the value associated with the key in the NaiveMap.
func (f *NaiveMap) Get(key string) interface{} {
	f.lock.Lock()
	defer f.lock.Unlock()

	// Goes through keys list to find key match
	for i, _ := range f.keys {
		if f.keys[i] == key {
			return f.values[i]
		}
	}
	return nil
}

// Delete removes the Value in the <Key, Value> pair from the hash map and returns the object, or nil otherwise.
func (f *NaiveMap) Delete(key string) interface{} {
	f.lock.Lock()
	defer f.lock.Unlock()

	// Kind match.
	for i, _ := range f.keys {
		if f.keys[i] == "" {
			return nil
		}
		//Shift all elements.
		if f.keys[i] == key {
			value := f.values[i]
			for k := i; k < len(f.keys)-1; k++ {
				if f.keys[k+1] == "" {
					f.keys[k] = ""
					f.values[k] = nil
					f.count--
					return value
				}
				f.keys[k] = f.keys[k+1]
				f.values[k] = f.values[k+1]
			}
			// Set the last element to nil.
			f.keys[len(f.keys)-1] = ""
			f.values[len(f.keys)-1] = nil
			f.count--
			return value
		}
	}
	log.Print("Delete function failure")
	return nil
}

// Load calculates the percentage of the NaiveMap that is full.
func (f *NaiveMap) Load() float64 {
	f.lock.Lock()
	defer f.lock.Unlock()

	return float64(f.count) / float64(f.capacity)
}
