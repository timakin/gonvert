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

func transformEncodingWithMeditation(rawReader *bytes.Reader, decoder transform.Transformer, encoder transform.Transformer) (string, error) {
	decoded, err := ioutil.ReadAll(transform.NewReader(rawReader, decoder))
	if err == nil {
		decodedBytes := []byte(decoded)
		encoded, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(decodedBytes), encoder))
		if err == nil {
			return string(encoded), nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}
