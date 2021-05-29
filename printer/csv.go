package printer

import (
	"io"

	"github.com/gocarina/gocsv"
	"github.com/hk220/go-circle-list-extract/circle"
)

func CSV(w io.Writer, cl *circle.Circles) error {
	err := gocsv.Marshal(*cl, w)
	if err != nil {
		return err
	}
	return nil
}
