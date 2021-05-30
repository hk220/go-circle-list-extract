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
package printer

import (
	"io"

	"github.com/hk220/go-circle-list-extract/circle"
)

type Printer func(w io.Writer, cl *circle.Circles) error

func (p Printer) Print(w io.Writer, cl *circle.Circles) error {
	if err := p(w, cl); err != nil {
		return err
	}
	return nil
}

var printer map[string]Printer = map[string]Printer{
	"json": JSON,
	"csv":  CSV,
}

func GetPrinter(s string) Printer {
	return printer[s]
}

func HasPrinter(s string) bool {
	_, ok := printer[s]
	return ok
}
