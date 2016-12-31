package gonvert

import (
	"github.com/k0kubun/pp"
	"github.com/saintfish/chardet"
	c "github.com/timakin/gonvert/converter"
)

type CharCode int

const (
	UTF8 CharCode = iota
	SJIS
)

var charcodes = map[string]CharCode{
	"UTF_8":     UTF8,
	"Shift_JIS": SJIS,
}

func Convert(text string, toCode CharCode) (result string, err error) {
	converter := createConverter(text, toCode)
	result, err = converter.Convert()
	return
}

func createConverter(text string, toCode CharCode) c.Converter {
	charDetector := chardet.NewTextDetector()
	detectResult, err := charDetector.DetectBest([]byte(text))
	if err != nil {
		pp.Fatal(err)
	}
	fromCode := charcodes[detectResult.Charset]
	var converter c.Converter

	if fromCode != UTF8 {
		if fromCode == SJIS {
			converter = &c.SjisToUTF8Converter{Text: text}
		} else {
			panic("cannot decode characters")
		}
	} else {
		converter = &c.DefaultConverter{Text: text}
	}

	return converter
}
