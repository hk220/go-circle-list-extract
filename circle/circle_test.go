package circle

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	expect := &Circles{
		{Space: "A01a", Name: "hoge", URL: "example.com"},
		{Space: "A02a", Name: "hoge", URL: ""},
	}
	actual := &Circles{
		{Space: "A01a", Name: "hoge", URL: "example.com"},
	}
	actual.Add(&Circle{Space: "A02a", Name: "hoge", URL: ""})

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Expect is not equal actual: %v", *actual)
	}
}

func TestString(t *testing.T) {
	expect := [][]string{
		{"A01a", "hoge", "example.com"},
		{"A02a", "hoge", ""},
	}
	cl := &Circles{
		{Space: "A01a", Name: "hoge", URL: "example.com"},
		{Space: "A02a", Name: "hoge", URL: ""},
	}
	actual := cl.String()
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("Expect is not equal actual: %v", actual)
	}
}
