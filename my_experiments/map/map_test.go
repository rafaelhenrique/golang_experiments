package testmap

import "testing"

func TestSomeFunction(t *testing.T) {
	result := someFunction("abc")
	expected := ""
	if result != expected {
		t.Errorf("result = %s, expected = %s.", result, expected)
	}

	result = someFunction("test")
	expected = "value"
	if result != expected {
		t.Errorf("result = %s, expected = %s.", result, expected)
	}
}
