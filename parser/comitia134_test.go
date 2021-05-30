package parser

import (
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestComitia134Parser(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(parserReader)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := Comitia134Parser(doc)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(parserExpected, *actual) {
		t.Errorf("Not match circles, expect: %+v, actual: %+v", parserExpected, actual)
	}
}
