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
	EUCJP
)

var charcodes = map[string]CharCode{
	"UTF_8":     UTF8,
	"Shift_JIS": SJIS,
	"EUC_JP":    EUCJP,
}

type CodePair struct {
	From, To CharCode
}

func New(text string, toCode CharCode) c.Converter {
	converter := createConverter(text, toCode)
	return converter
}

func createConverter(text string, toCode CharCode) c.Converter {
	codeConverters := make(map[CodePair]string)
	codeConverters[CodePair{SJIS, UTF8}] = "SJISToUTF8Converter"
	codeConverters[CodePair{UTF8, SJIS}] = "UTF8ToSJISConverter"
	codeConverters[CodePair{EUCJP, UTF8}] = "EUCJPToUTF8Converter"
	codeConverters[CodePair{UTF8, EUCJP}] = "UTF8ToEUCJPConverter"

	charDetector := chardet.NewTextDetector()
	detectResult, err := charDetector.DetectBest([]byte(text))
	if err != nil {
		pp.Fatal(err)
	}
	fromCode := charcodes[detectResult.Charset]
	var converter c.Converter

	if fromCode == toCode {
		converter = &c.DefaultConverter{Text: text}
	} else {
		converterStr := codeConverters[CodePair{fromCode, toCode}]
		pp.Print(converterStr)
		converter = &c.DefaultConverter{Text: text}
	}

	return converter
}
