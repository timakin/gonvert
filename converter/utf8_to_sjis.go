package converter

import (
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

type UTF8ToSjisConverter ConversionPattern

func (c *UTF8ToSjisConverter) Convert() (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(c.Text), japanese.ShiftJIS.NewEncoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}
