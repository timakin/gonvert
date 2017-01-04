package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
)

type SJISToUTF8Converter ConversionPattern

func (c *SJISToUTF8Converter) Convert() (string, error) {
	return transformEncoding(bytes.NewReader(c.TextByte), japanese.ShiftJIS.NewDecoder())
}
