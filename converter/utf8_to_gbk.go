package converter

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

type UTF8ToGBKConverter ConversionPattern

func (c *UTF8ToGBKConverter) Convert() (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(c.TextByte), simplifiedchinese.GBK.NewEncoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}
