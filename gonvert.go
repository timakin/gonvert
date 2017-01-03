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
	charDetector := chardet.NewTextDetector()
	textByte := []byte(text)
	detectResult, err := charDetector.DetectBest(textByte)
	if err != nil {
		pp.Fatal(err)
	}
	fromCode := charcodes[detectResult.Charset]
	pp.Print(fromCode)
	var converter c.Converter

	codepair := CodePair{fromCode, toCode}

	switch codepair {
	case (CodePair{UTF8, SJIS}):
		{
			converter = &c.UTF8ToSJISConverter{TextByte: textByte}
		}
	case (CodePair{SJIS, UTF8}):
		{
			converter = &c.SJISToUTF8Converter{TextByte: textByte}
		}
	case (CodePair{UTF8, EUCJP}):
		{
			converter = &c.UTF8ToEUCJPConverter{TextByte: textByte}
		}
	case (CodePair{EUCJP, UTF8}):
		{
			converter = &c.EUCJPToUTF8Converter{TextByte: textByte}
		}
	default:
		{
			converter = &c.DefaultConverter{TextByte: textByte}
		}
	}
	return converter
}
