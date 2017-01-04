package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
)

type EUCJPToUTF8Converter ConversionPattern

func (c *EUCJPToUTF8Converter) Convert() (string, error) {
	return transformEncoding(bytes.NewReader(c.TextByte), japanese.EUCJP.NewDecoder())
}
