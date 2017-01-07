package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type SJISToGBKConverter ConversionPattern

func (c *SJISToGBKConverter) Convert() (string, error) {
	return transformEncodingWithMeditation(bytes.NewReader(c.TextByte), japanese.ShiftJIS.NewDecoder(), simplifiedchinese.GBK.NewEncoder())
}
