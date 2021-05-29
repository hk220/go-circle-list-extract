package printer

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hk220/go-circle-list-extract/circle"
)

func TestCSV(t *testing.T) {
	cl := &circle.CircleList{
		Circles: []circle.Circle{
			{Space: "A01a", Name: "foo", URL: "example.com"},
			{Space: "A02a", Name: "bar", URL: ""},
		},
	}

	var p Printer = CSV

	actual := &bytes.Buffer{}

	expect := []byte(`space,name,url
A01a,foo,example.com
A02a,bar,
`)

	p.Print(actual, cl)

	if !reflect.DeepEqual(expect, actual.Bytes()) {
		t.Errorf("Expect is not equal actual: %v", actual)
	}

}
