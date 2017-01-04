package converter

import (
	"bytes"
	"golang.org/x/text/transform"
	"io/ioutil"
)

type Converter interface {
	Convert() (string, error)
}

type ConversionPattern struct {
	TextByte []byte
}

func transformEncoding(rawReader *bytes.Reader, trans transform.Transformer) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err == nil {
		return string(ret), nil
	} else {
		return "", err
	}
}
