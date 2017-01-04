package converter

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type GBKToUTF8Converter ConversionPattern

func (c *GBKToUTF8Converter) Convert() (string, error) {
	return transformEncoding(bytes.NewReader(c.TextByte), simplifiedchinese.GBK.NewDecoder())
}
