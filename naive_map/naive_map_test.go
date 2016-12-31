// Test functions in the FixedHashMap.
package NaiveMap

import "testing"

func TestEmptyMap(t *testing.T) {
	t.Log("Empty Naive")
	game := NewNaiveMap(100)
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
	main := NewNaiveMap(10)
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
	sample := NewNaiveMap(2)
	sample.Set("rock", 1)
	sample.Set("help", 3)
	answer := sample.Set("ew", "nope")
	if answer != false {
		t.Error("Set should have failed due to over capacity")
	}
}
