package converter

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

type GBKToUTF8Converter ConversionPattern

func (c *GBKToUTF8Converter) Convert() (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(c.TextByte), simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}
