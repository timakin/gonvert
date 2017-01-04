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
	GBK
)

var charcodes = map[string]CharCode{
	"UTF_8":     UTF8,
	"Shift_JIS": SJIS,
	"EUC-JP":    EUCJP,
}

type CodePair struct {
	From, To CharCode
}

func New(text string, codes ...CharCode) c.Converter {
	converter := createConverter(text, codes)

	return converter
}

func createConverter(text string, codes []CharCode) c.Converter {
	textByte := []byte(text)
	var fromCode, toCode CharCode

	if len(codes) == 1 {
		charDetector := chardet.NewTextDetector()
		detectResult, err := charDetector.DetectBest(textByte)
		if err != nil {
			pp.Fatal(err)
		}
		fromCode = charcodes[detectResult.Charset]
		toCode = codes[0]
	} else if len(codes) >= 2 {
		toCode = codes[0]
		fromCode = codes[1]
	} else {
		panic("Please specify 1 or 2 character-code arguements.")
	}

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
	case (CodePair{UTF8, GBK}):
		{
			converter = &c.UTF8ToGBKConverter{TextByte: textByte}
		}
	case (CodePair{GBK, UTF8}):
		{
			converter = &c.GBKToUTF8Converter{TextByte: textByte}
		}
	default:
		{
			converter = &c.DefaultConverter{TextByte: textByte}
		}
	}
	return converter
}
