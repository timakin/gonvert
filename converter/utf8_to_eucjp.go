package converter

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

type UTF8ToEUCJPConverter ConversionPattern

func (c *UTF8ToEUCJPConverter) Convert() (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(c.Text), japanese.EUCJP.NewEncoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}
