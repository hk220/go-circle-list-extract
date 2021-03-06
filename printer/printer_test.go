package printer

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/hk220/go-circle-list-extract/circle"
)

func TestPrint(t *testing.T) {
	cl := &circle.Circles{
		{Space: "A01a", Name: "foo", URL: "example.com"},
		{Space: "A02a", Name: "bar", URL: ""},
	}

	var p Printer = func(w io.Writer, cl *circle.Circles) error {
		for _, c := range *cl {
			fmt.Fprintln(w, c.Name)
		}
		return nil
	}

	actual := &bytes.Buffer{}

	expect := []byte(`foo
bar
`)

	p.Print(actual, cl)

	if !reflect.DeepEqual(expect, actual.Bytes()) {
		t.Errorf("Expect is not equal actual: %v", actual)
	}

}
