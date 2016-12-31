// Package FixedMap tests.
package FixedHashMap

import (
	"log"
	"testing"
)

func TestHash(t *testing.T) {
	var input map[string]int
	input = make(map[string]int)
	size := 500
	input["Help"] = hash("Help", size)
	input["help"] = hash("help", size)
	input["KPCB"] = hash("KPCB", size)
	input["Fellow"] = hash("Fellow", size)
	input["2017"] = hash("2017", size)
	for k, v := range input {
		log.Println(v)
		if v > size || size < 0 {
			t.Errorf("Invalid hash, got %v for %d", v, k)
		}
	}
}

func TestEmptyMap(t *testing.T) {
	t.Log("Empty Naive")
	game := NewFixedHashMap(100)
	check := game.Get("Hot")
	help := game.Delete("hot")

	if score := game.Load(); score != 0.0 {
		t.Errorf("Expected load of 0, but it was %d instead.", score)
	}
	if check != nil {
		t.Errorf("Expected check to be nil, got %v.", check)
	}
	if help != nil {
		t.Errorf("Expected help to be nil, got %v.", help)
	}
}

func TestBaseTest(t *testing.T) {
	t.Log("Base Test")
	main := NewFixedHashMap(10)
	main.Set("First", 1)
	main.Set("First", "Geat")
	main.Set("Second", 2)
	main.Set("Thrid", true)

	interest := main.Get("First")
	if interest != "Geat" {
		t.Errorf("Expected Geat, but it was %d instead.", interest)
	}

	main.Delete("First")
	interest = main.Get("First")
	if interest == "Geat" {
		t.Errorf("Expected nil, but it was %d instead.", interest)
	}

	load := main.Load()
	if load != 0.2 {
		t.Errorf("Expected .2, but it was %v instead.", load)
	}
}

func TestCapacity(t *testing.T) {
	t.Log("Capacity Test")
	sample := NewFixedHashMap(2)
	sample.Set("rock", 1)
	sample.Set("help", 3)
	answer := sample.Set("ew", "nope")
	if answer != false {
		t.Error("Set should have failed due to over capacity")
	}
}
