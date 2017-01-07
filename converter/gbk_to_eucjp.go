package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type GBKToEUCJPConverter ConversionPattern

func (c *GBKToEUCJPConverter) Convert() (string, error) {
	return transformEncodingWithMeditation(bytes.NewReader(c.TextByte), simplifiedchinese.GBK.NewDecoder(), japanese.EUCJP.NewEncoder())
}
