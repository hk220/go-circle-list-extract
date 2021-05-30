package printer

import (
	"encoding/json"
	"io"

	"github.com/hk220/go-circle-list-extract/circle"
)

func JSON(w io.Writer, cl *circle.Circles) error {
	s, err := json.MarshalIndent(*cl, "", "\t")
	if err != nil {
		return err
	}

	_, err = w.Write(s)
	if err != nil {
		return err
	}

	return nil
}
