package gonvert

import (
	"github.com/saintfish/chardet"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = []struct {
	Before CharCode
	After  CharCode
	Val    string
}{
	{SJIS, UTF8, "\x82\xb1\x82\xea\x82\xcd\x8a\xbf\x8e\x9a\x82\xc5\x82\xb7\x81B"},
	{UTF8, SJIS, "これは漢字です"},
}

func TestSjisToUTF8Convert(t *testing.T) {
	converter := New("\x82\xb1\x82\xea\x82\xcd\x8a\xbf\x8e\x9a\x82\xc5\x82\xb7\x81B", UTF8)
	result, err := converter.Convert()
	if err != nil {
		panic(err)
	}
	charDetector := chardet.NewTextDetector()
	detectResult, err := charDetector.DetectBest([]byte(result))
	if err != nil {
		panic(err)
	}
	code := charcodes[detectResult.Charset]

	t.Log(result)
	t.Log(detectResult.Charset)
	assert.Equal(t, UTF8, code, "SJIS to UTF8")
}

func TestUTF8ToSjisConvert(t *testing.T) {
	converter := New("これは漢字です", SJIS)
	result, err := converter.Convert()
	if err != nil {
		panic(err)
	}
	charDetector := chardet.NewTextDetector()
	detectResult, err := charDetector.DetectBest([]byte(result))
	if err != nil {
		panic(err)
	}
	code := charcodes[detectResult.Charset]

	t.Log(result)
	t.Log(detectResult.Charset)
	assert.Equal(t, SJIS, code, "UTF8 to SJIS")
}

func TestConvert(t *testing.T) {
	for _, data := range testData {
		converter := New(data.Val, data.After)

		charDetector := chardet.NewTextDetector()
		beforeDetect, err := charDetector.DetectBest([]byte(data.Val))
		if err != nil {
			panic(err)
		}
		assert.Equal(t, beforeDetect.Charset, data.Before, "before convert")

		result, err := converter.Convert()
		if err != nil {
			panic(err)
		}
		afterDetect, err := charDetector.DetectBest([]byte(result))
		if err != nil {
			panic(err)
		}
		assert.Equal(t, afterDetect.Charset, data.After, "after convert")
	}
}
