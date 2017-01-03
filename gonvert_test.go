package gonvert

import (
	"github.com/saintfish/chardet"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = []struct {
	BeforeCode, AfterCode CharCode
	BeforeText, AfterText string
}{
	{SJIS, UTF8, "\x8c\x8e\x93\xfa\x82\xcd\x95\x53\x91\xe3\x82\xcc\x89\xdf\x8b\x71\x82\xc9\x82\xb5\x82\xc4\x81\x41\x8d\x73\x82\xa9\x82\xd3\x94\x4e\x82\xe0\x96\x94\x97\xb7\x90\x6c\x96\xe7\x81\x42", "月日は百代の過客にして、行かふ年も又旅人也。"},
	{UTF8, SJIS, "月日は百代の過客にして、行かふ年も又旅人也。", "\x8c\x8e\x93\xfa\x82\xcd\x95\x53\x91\xe3\x82\xcc\x89\xdf\x8b\x71\x82\xc9\x82\xb5\x82\xc4\x81\x41\x8d\x73\x82\xa9\x82\xd3\x94\x4e\x82\xe0\x96\x94\x97\xb7\x90\x6c\x96\xe7\x81\x42"},
	{EUCJP, UTF8, "\xa4\xb3\xa4\xf3\xa4\xcb\xa4\xc1\xa4\xcf\xa1\xa2Python\xa5\xd7\xa5\xed\xa5\xb0\xa5\xe9\xa5\xdf\xa5\xf3\xa5\xb0", "こんにちは、Pythonプログラミング"},
	{UTF8, EUCJP, "こんにちは、Pythonプログラミング", "\xa4\xb3\xa4\xf3\xa4\xcb\xa4\xc1\xa4\xcf\xa1\xa2Python\xa5\xd7\xa5\xed\xa5\xb0\xa5\xe9\xa5\xdf\xa5\xf3\xa5\xb0"},
}

func TestConvert(t *testing.T) {
	for _, data := range testData {
		converter := New(data.BeforeText, data.AfterCode)

		charDetector := chardet.NewTextDetector()
		beforeDetect, err := charDetector.DetectBest([]byte(data.BeforeText))
		if err != nil {
			panic(err)
		}
		assert.Equal(t, charcodes[beforeDetect.Charset], data.BeforeCode, "before convert")

		result, err := converter.Convert()
		if err != nil {
			panic(err)
		}
		afterDetect, err := charDetector.DetectBest([]byte(result))
		if err != nil {
			panic(err)
		}

		assert.Equal(t, charcodes[afterDetect.Charset], data.AfterCode, "after convert")
		assert.Equal(t, result, data.AfterText, "after convert")
	}
}
