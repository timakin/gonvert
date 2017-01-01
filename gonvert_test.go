package gonvert

import (
	"github.com/saintfish/chardet"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSjisToUTF8Convert(t *testing.T) {
	converter := New("\x89\xD4\x8e\x71", UTF8)
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
	t.Log(detectResult.Charset)
	assert.Equal(t, UTF8, code, "SJIS to UTF8")
}

func TestUTF8ToSjisConvert(t *testing.T) {
	converter := New("花子", SJIS)
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
	t.Log(detectResult.Charset)
	assert.Equal(t, SJIS, code, "UTF8 to SJIS")
}
