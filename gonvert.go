package gonvert

import (
	"github.com/k0kubun/pp"
	"github.com/saintfish/chardet"
	c "github.com/timakin/gonvert/converter"
	"reflect"
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

var converterCache = make(map[string]reflect.Type)

var codeConverters = make(map[CodePair]string)

func New(text string, toCode CharCode) c.Converter {
	converter := createConverter(text, toCode)
	return converter
}

func prepare() {
	codeConverters[CodePair{SJIS, UTF8}] = "github.com/timakin/gonvert/converter.SJISToUTF8Converter"
	codeConverters[CodePair{UTF8, SJIS}] = "github.com/timakin/gonvert/converter.UTF8ToSJISConverter"
	codeConverters[CodePair{EUCJP, UTF8}] = "github.com/timakin/gonvert/converter.EUCJPToUTF8Converter"
	codeConverters[CodePair{UTF8, EUCJP}] = "github.com/timakin/gonvert/converter.UTF8ToEUCJPConverter"

	registerConverterCache(c.SJISToUTF8Converter{})
	registerConverterCache(c.UTF8ToSJISConverter{})
	registerConverterCache(c.EUCJPToUTF8Converter{})
	registerConverterCache(c.UTF8ToEUCJPConverter{})
}

func registerConverterCache(c interface{}) {
	t := reflect.TypeOf(c)
	converterCache[t.PkgPath()+"."+t.Name()] = t
}

func createConverter(text string, toCode CharCode) c.Converter {
	prepare()
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
		t, ok := converterCache[converterStr]
		if !ok {
			panic("Reflect error")
		}
		pp.Print(t)
		v := reflect.New(t).Interface()
		pp.Print(v)
		converter = &c.DefaultConverter{Text: text}
	}

	return converter
}
