// Package FixedHashMap is an implementation of a fixed hash-map under the specifications outlied @  kpcbfellows.com/engineering/apply.
package FixedHashMap

import (
	"log"
	"sync"
)

// Value to scale used for hashing.
const MULTIPLICATION_FACTOR = 100

// bucket contains actual value of key and value and linked together with same hash values
type bucket struct {
	// Thread uniqueness safeguard.
	lock sync.Mutex
	// Value of the bucket.
	value interface{}
	// Original input key.
	key string
	// Next bucket with same hash.
	next *bucket
}

// newBucket is a constructor for bucket.
func newBucket(value interface{}, key string) *bucket {
	initial := &bucket{
		sync.Mutex{},
		value,
		key,
		nil,
	}
	return initial
}

// Hashing inspired from http://www.manniwood.com/, but modified.
func hash(s string, size int) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = 31*h + int(s[i])
	}
	value := (h & 0x7fffffff) % size
	if value == 0 {
		value++
	}
	return value
}

// FixedHashMap is the struct initialization of the fixed size hash map.
type FixedHashMap struct {
	// For thread safety
	lock sync.Mutex
	// Array of keys that are the first ones to show.
	keys []string
	// Array of object values assigned to keys by index for the map.
	values []*bucket
	// Maximium number of <key, value> pairs possible to fit in map.
	capacity int
	// Number of elements currently in map
	count int
}

// NewFixedHashMap creates the FixedSizeHashMap using the input construction size.
func NewFixedHashMap(size int) *FixedHashMap {
	initial := &FixedHashMap{
		sync.Mutex{},
		make([]string, size*MULTIPLICATION_FACTOR),
		make([]*bucket, size*MULTIPLICATION_FACTOR),
		size,
		0,
	}
	return initial

}

// Set stores the <Key, Value> pair in the FixedHashMap and returns a boolean if the operation was successful.
func (f *FixedHashMap) Set(key string, value interface{}) bool {
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

	hashValue := hash(key, f.capacity*MULTIPLICATION_FACTOR)
	if f.keys[hashValue] == "" && f.values[hashValue] == nil {
		f.keys[hashValue] = key
		f.values[hashValue] = newBucket(value, key)
		f.count++
		return true
	}
	if f.keys[hashValue] == key {
		f.values[hashValue].value = value
		return true
	}
	for e := f.values[hashValue]; e != nil; e = e.next {
		if e.key == key {
			e.value = value
			return true
		}
	}
	return false
}

// Get returns the value associated with the key in the FixedHashMap.
func (f *FixedHashMap) Get(key string) interface{} {
	f.lock.Lock()
	defer f.lock.Unlock()

	return nil
}

// Delete removes the Value in the <Key, Value> pair from the hash map and returns the object, or nil otherwise.
func (f *FixedHashMap) Delete(key string) interface{} {
	f.lock.Lock()
	defer f.lock.Unlock()

	return nil
}

// Load calculates the percentage of the FixedHashMap that is full.
func (f *FixedHashMap) Load() float64 {
	f.lock.Lock()
	defer f.lock.Unlock()

	return float64(f.count / f.capacity)
}
