package value

import "testing"

func TestIntToInt(t *testing.T) {
	v := IntValue{value: 1}
	i := v.ToInt()

	if i != 1 {
		t.Fatalf(`Failed to cast to int: %v %v`, v, i)
	}
}

func TestIntToFloat(t *testing.T) {
	v := IntValue{value: 1}
	i := v.ToFloat()

	if i != 1.0 {
		t.Fatalf(`Failed to cast to int: %v %v`, v, i)
	}
}

func TestIntToString(t *testing.T) {
	v := IntValue{value: 1}
	i := v.ToString()

	if i != "1" {
		t.Fatalf(`Failed to cast to int: %v %v`, v, i)
	}
}

func TestIntLessThan(t *testing.T) {
	v := IntValue{value: 1}
	v2 := StringValue{value: "2"}

	if !v.LessThan(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestIntLessThanEqual(t *testing.T) {
	v := IntValue{value: 1}
	v2 := StringValue{value: "1"}

	if !v.LessThanEqual(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestIntGreaterThan(t *testing.T) {
	v := IntValue{value: 2}
	v2 := StringValue{value: "1"}

	if !v.GreaterThan(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestIntGreaterThanEqual(t *testing.T) {
	v := IntValue{value: 2}
	v2 := StringValue{value: "2"}

	if !v.GreaterThanEqual(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestIntEqual(t *testing.T) {
	v := IntValue{value: 2}
	v2 := StringValue{value: "2"}

	if !v.Equal(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}

func TestIntDifferent(t *testing.T) {
	v := IntValue{value: 2}
	v2 := StringValue{value: "2.1"}

	if !v.Different(v2) {
		t.Fatalf(`Failed to compare %v %v`, v, v2)
	}
}
