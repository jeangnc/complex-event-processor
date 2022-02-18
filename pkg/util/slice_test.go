package util

import (
	"testing"
)

func TestSliceAll(t *testing.T) {
	if SliceAll([]bool{true, false}) {
		t.Fatalf("Failed to assert that not all elements are true")
	}

	if !SliceAll([]bool{true, true}) {
		t.Fatalf("Failed to assert that all elements are true")
	}
}

func TestSliceAny(t *testing.T) {
	if !SliceAny([]bool{true, false}) {
		t.Fatalf("Failed to assert that some elements are true")
	}

	if SliceAny([]bool{false, false}) {
		t.Fatalf("Failed to assert that all elements are false")
	}
}
