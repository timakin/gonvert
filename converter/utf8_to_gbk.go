package converter

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type UTF8ToGBKConverter ConversionPattern

func (c *UTF8ToGBKConverter) Convert() (string, error) {
	return transformEncoding(bytes.NewReader(c.TextByte), simplifiedchinese.GBK.NewEncoder())
}
