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
