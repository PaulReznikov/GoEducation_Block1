package main

import (
	"reflect"
	"testing"
)

func TestEfficiency(t *testing.T) {
	words := []string{"listen", "silent", "enlist", "inlets", "google", "gogole"}
	expected := map[string][]string{"eilnst": {"listen", "silent", "enlist", "inlets"}, "eggloo": {"google", "gogole"}}
	actual := GroupAnagrams(words)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Result was incorrect,\n got: %s,\n want: %s.", actual, expected)
	}

}

func TestEqualKeys(t *testing.T) {
	words := []string{"listen", "silent", "enlist", "inlets", "google", "gogole"}
	expected := map[string][]string{"eilnst": {"listen", "silent", "enlist", "inlets"}, "eggloo": {"google", "gogole"}}
	actual := GroupAnagrams(words)

	for k := range expected {
		if _, ok := actual[k]; !ok {
			t.Errorf("Result was incorrect,\n got: %s,\n want: %s.", actual, expected)
		}
	}
}

func TestEqualLenValues(t *testing.T) {
	words := []string{"listen", "silent", "enlist", "inlets", "google", "gogole", "silent", "enlist", "google"}
	expected := map[string][]string{"eilnst": {"listen", "silent", "enlist", "inlets"}, "eggloo": {"google", "gogole"}}
	actual := GroupAnagrams(words)

	for k := range expected {
		if len(expected[k]) != len(actual[k]) {
			t.Errorf("Result was incorrect,\n got: %s,\n want: %s.", actual, expected)
		}
	}
}
