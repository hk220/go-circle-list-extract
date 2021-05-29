package printer

import (
	"io"

	"github.com/hk220/go-circle-list-extract/circle"
)

type Printer func(w io.Writer, cl *circle.CircleList) error

func (p Printer) Print(w io.Writer, cl *circle.CircleList) error {
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
