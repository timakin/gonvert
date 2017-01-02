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
	detectResult, err := charDetector.DetectBest([]byte(text))
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
			converter = &c.UTF8ToSJISConverter{Text: text}
		}
	case (CodePair{SJIS, UTF8}):
		{
			converter = &c.SJISToUTF8Converter{Text: text}
		}
	case (CodePair{UTF8, EUCJP}):
		{
			converter = &c.UTF8ToEUCJPConverter{Text: text}
		}
	case (CodePair{EUCJP, UTF8}):
		{
			converter = &c.EUCJPToUTF8Converter{Text: text}
		}
	default:
		{
			converter = &c.DefaultConverter{Text: text}
		}
	}
	pp.Print(converter)

	return converter
}
