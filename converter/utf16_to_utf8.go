package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
)

type UTF16ToUTF8Converter ConversionPattern

func (c *UTF16ToUTF8Converter) Convert() (string, error) {
	return transformEncoding(bytes.NewReader(c.TextByte), japanese.ShiftJIS.NewEncoder())
}
