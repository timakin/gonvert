package converter

import (
	"bytes"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

type UTF8ToSJISConverter ConversionPattern

func (c *UTF8ToSJISConverter) Convert() (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(c.TextByte), japanese.ShiftJIS.NewEncoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}
