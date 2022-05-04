package value

import (
	"testing"
)

func TestFloatToInt(t *testing.T) {
	v := FloatValue{value: 1.0}
	i := v.ToInt()

	if i != 1 {
		t.Fatalf(`Failed to cast to int: %v %v`, v, i)
	}
}

func TestFloatToFloat(t *testing.T) {
	v := FloatValue{value: 1.0}
	i := v.ToFloat()

	if i != 1.0 {
		t.Fatalf(`Failed to cast to int: %v %v`, v, i)
	}
}

func TestFloatToString(t *testing.T) {
	v := FloatValue{value: 1.0}
	i := v.ToString()

	if i != "1" {
		t.Fatalf(`Failed to cast to int: %v %v`, v, i)
	}
}

func TestFloatLessThan(t *testing.T) {
	v := FloatValue{value: 1.0}
	v2 := StringValue{value: "2"}

	if !v.LessThan(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestFloatLessThanEqual(t *testing.T) {
	v := FloatValue{value: 1.0}
	v2 := StringValue{value: "1"}

	if !v.LessThanEqual(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestFloatGreaterThan(t *testing.T) {
	v := FloatValue{value: 2.0}
	v2 := StringValue{value: "1"}

	if !v.GreaterThan(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestFloatGreaterThanEqual(t *testing.T) {
	v := FloatValue{value: 2.0}
	v2 := StringValue{value: "2"}

	if !v.GreaterThanEqual(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestFloatEqual(t *testing.T) {
	v := FloatValue{value: 2.0}
	v2 := StringValue{value: "2"}

	if !v.Equal(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestFloatDifferent(t *testing.T) {
	v := FloatValue{value: 2.0}
	v2 := StringValue{value: "2.1"}

	if !v.Different(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}
