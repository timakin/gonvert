package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
)

type EUCJPToSJISConverter ConversionPattern

func (c *EUCJPToSJISConverter) Convert() (string, error) {
	return transformEncodingWithMeditation(bytes.NewReader(c.TextByte), japanese.EUCJP.NewDecoder(), japanese.ShiftJIS.NewEncoder())
}
