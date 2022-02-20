package util

import (
	"reflect"
	"testing"
)

// Tests whether a set of changes impacted the expression
func TestMapMerge(t *testing.T) {
	mapA := map[string]bool{
		"a": true,
	}
	mapB := map[string]bool{
		"b": true,
	}
	expectedResult := map[string]bool{
		"a": true,
		"b": true,
	}

	result, changes := MergeMaps(mapA, mapB)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Fatalf(`Failed to merge %v %v. Result : %v`, mapA, mapB, result)
	}

	expectedChanges := []string{"b"}
	if !reflect.DeepEqual(changes, expectedChanges) {
		t.Fatalf(`Failed to assert changes %v %v. Changes : %v`, mapA, mapB, changes)
	}
}

// Tests map copying
func TestMapCopy(t *testing.T) {
	myMap := map[string]string{
		"key": "value",
	}

	copiedMap := CopyMap(myMap)
	if !reflect.DeepEqual(myMap, copiedMap) {
		t.Fatalf(`Failed to copy map %v %v`, myMap, copiedMap)
	}
}

// Tests map reference
func TestCopiedMapRefere(t *testing.T) {
	myMap := map[string]string{
		"key": "value",
	}

	copiedMap := CopyMap(myMap)
	copiedMap["another-key"] = "another-value"

	if reflect.DeepEqual(myMap, copiedMap) {
		t.Fatalf(`Original map mutated %v %v`, myMap, copiedMap)
	}
}
