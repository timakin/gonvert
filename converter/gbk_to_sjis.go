package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type GBKToSJISConverter ConversionPattern

func (c *GBKToSJISConverter) Convert() (string, error) {
	return transformEncodingWithMeditation(bytes.NewReader(c.TextByte), simplifiedchinese.GBK.NewDecoder(), japanese.ShiftJIS.NewEncoder())
}
