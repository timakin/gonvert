package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
)

type UTF8ToSJISConverter ConversionPattern

func (c *UTF8ToSJISConverter) Convert() (string, error) {
	return transformEncoding(bytes.NewReader(c.TextByte), japanese.ShiftJIS.NewEncoder())
}
