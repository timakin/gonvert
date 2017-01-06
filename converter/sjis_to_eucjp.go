package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
)

type SJISToEUCJPConverter ConversionPattern

func (c *SJISToEUCJPConverter) Convert() (string, error) {
	return transformEncodingWithMeditation(bytes.NewReader(c.TextByte), japanese.ShiftJIS.NewDecoder(), japanese.EUCJP.NewEncoder())
}
