package printer

import (
	"encoding/json"
	"io"

	"github.com/hk220/go-circle-list-extract/circle"
)

func JSON(w io.Writer, cl *circle.CircleList) error {
	s, err := json.MarshalIndent(cl.Circles, "", "\t")
	if err != nil {
		return err
	}

	_, err = w.Write(s)
	if err != nil {
		return err
	}

	return nil
}
