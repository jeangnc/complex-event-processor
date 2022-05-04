package value

import "testing"

func TestStringToInt(t *testing.T) {
	v := StringValue{value: "1"}
	i := v.ToInt()

	if i != 1 {
		t.Fatalf(`Failed to cast to int: %v %v`, v, i)
	}
}

func TestStringToFloat(t *testing.T) {
	v := StringValue{value: "1"}
	i := v.ToFloat()

	if i != 1.0 {
		t.Fatalf(`Failed to cast to int: %v %v`, v, i)
	}
}

func TestStringToString(t *testing.T) {
	v := StringValue{value: "1"}
	i := v.ToString()

	if i != "1" {
		t.Fatalf(`Failed to cast to int: %v %v`, v, i)
	}
}

func TestStringLessThan(t *testing.T) {
	v := StringValue{value: "1"}
	v2 := StringValue{value: "2"}

	if !v.LessThan(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestStringLessThanEqual(t *testing.T) {
	v := StringValue{value: "1"}
	v2 := StringValue{value: "1"}

	if !v.LessThanEqual(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestStringGreaterThan(t *testing.T) {
	v := StringValue{value: "2"}
	v2 := StringValue{value: "1"}

	if !v.GreaterThan(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestStringGreaterThanEqual(t *testing.T) {
	v := StringValue{value: "2"}
	v2 := StringValue{value: "2"}

	if !v.GreaterThanEqual(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestStringEqual(t *testing.T) {
	v := StringValue{value: "2"}
	v2 := FloatValue{value: 2.0}

	if !v.Equal(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestStringDifferent(t *testing.T) {
	v := StringValue{value: "2"}
	v2 := FloatValue{value: 2.1}

	if !v.Different(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}
