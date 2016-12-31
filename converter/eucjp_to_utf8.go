package converter

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

type EUCJPToUTF8Converter ConversionPattern

func (c *EUCJPToUTF8Converter) Convert() (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(c.Text), japanese.EUCJP.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}
