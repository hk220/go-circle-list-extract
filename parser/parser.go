/*
MIT License

Copyright (c) 2021 Kazuki Hara

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hk220/go-circle-list-extract/circle"
	"golang.org/x/text/unicode/norm"
)

type Parser func(doc *goquery.Document) (*circle.Circles, error)

func (p Parser) Parse(doc *goquery.Document) (*circle.Circles, error) {
	cl, err := p(doc)
	if err != nil {
		return nil, err
	}
	return cl, nil
}

var parser map[string]Parser = map[string]Parser{
	"Comitia134Parser": Comitia134Parser,
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
