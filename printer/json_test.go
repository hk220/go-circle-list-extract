package printer

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hk220/go-circle-list-extract/circle"
)

func TestJSON(t *testing.T) {
	cl := &circle.Circles{
		{Space: "A01a", Name: "foo", URL: "example.com"},
		{Space: "A02a", Name: "bar", URL: ""},
	}

	var p Printer = JSON

	actual := &bytes.Buffer{}

	expect := []byte(`[
	{
		"space": "A01a",
		"name": "foo",
		"url": "example.com"
	},
	{
		"space": "A02a",
		"name": "bar",
		"url": ""
	}
]`)

	p.Print(actual, cl)

	if !reflect.DeepEqual(expect, actual.Bytes()) {
		t.Errorf("Expect is not equal actual: %v", actual)
	}

}
