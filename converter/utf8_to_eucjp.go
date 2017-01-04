package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
)

type UTF8ToEUCJPConverter ConversionPattern

func (c *UTF8ToEUCJPConverter) Convert() (string, error) {
	return transformEncoding(bytes.NewReader(c.TextByte), japanese.EUCJP.NewEncoder())
}
