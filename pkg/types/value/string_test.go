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
