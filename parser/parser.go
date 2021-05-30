/*
Copyright © 2021 Kazuki Hara

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
