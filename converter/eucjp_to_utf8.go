package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

type EUCJPToUTF8Converter ConversionPattern

func (c *EUCJPToUTF8Converter) Convert() (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(c.TextByte), japanese.EUCJP.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}
