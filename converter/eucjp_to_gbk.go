package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type EUCJPToGBKConverter ConversionPattern

func (c *EUCJPToGBKConverter) Convert() (string, error) {
	return transformEncodingWithMeditation(bytes.NewReader(c.TextByte), japanese.EUCJP.NewDecoder(), simplifiedchinese.GBK.NewEncoder())
}
