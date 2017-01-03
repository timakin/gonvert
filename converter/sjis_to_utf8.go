package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

type SJISToUTF8Converter ConversionPattern

func (c *SJISToUTF8Converter) Convert() (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(c.TextByte), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}
