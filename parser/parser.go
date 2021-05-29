package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hk220/go-circle-list-extract/circle"
	"golang.org/x/text/unicode/norm"
)

type Parser func(doc *goquery.Document) (*circle.CircleList, error)

func (p Parser) Parse(doc *goquery.Document) (*circle.CircleList, error) {
	cl, err := p(doc)
	if err != nil {
		return nil, err
	}
	return cl, nil
}

var parser map[string]Parser = map[string]Parser{
	"comitia134": Comitia134Parser,
}

func GetParser(s string) Parser {
	return parser[s]
}

func trimSpace(s string) string {
	return strings.TrimSpace(s)
}

func normalize(s string) string {
	return strings.TrimSpace(norm.NFKC.String(s))
}
