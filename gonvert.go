package gonvert

import (
	"github.com/k0kubun/pp"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

type Converter interface {
	Convert() (string, error)
}

type ConversionPattern struct {
	Text string
}

type CharCode int

const (
	UTF8 CharCode = iota
	SJIS
)

var charcodes = map[string]CharCode{
	"UTF_8":     UTF8,
	"Shift_JIS": SJIS,
}

type SjisToUTF8Converter ConversionPattern

func (c *SjisToUTF8Converter) Convert() (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(c.Text), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}

type PlainConverter ConversionPattern

func (c *PlainConverter) Convert() (string, error) {
	return c.Text, nil
}

func Convert(text string, toCode CharCode) (result string, err error) {
	converter := createConverter(text, toCode)
	result, err = converter.Convert()
	return
}

func createConverter(text string, toCode CharCode) Converter {
	charDetector := chardet.NewTextDetector()
	detectResult, err := charDetector.DetectBest([]byte(text))
	if err != nil {
		pp.Fatal(err)
	}
	fromCode := charcodes[detectResult.Charset]
	var converter Converter

	if fromCode != UTF8 {
		if fromCode == SJIS {
			converter = &SjisToUTF8Converter{Text: text}
		} else {
			panic("cannot decode characters")
		}
	} else {
		converter = &PlainConverter{Text: text}
	}

	return converter
}
