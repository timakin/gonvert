package gonvert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var estimationTestData = []struct {
	BeforeCode, AfterCode CharCode
	BeforeText, AfterText string
}{
	{SJIS, UTF8, "\x8c\x8e\x93\xfa\x82\xcd\x95\x53\x91\xe3\x82\xcc\x89\xdf\x8b\x71\x82\xc9\x82\xb5\x82\xc4\x81\x41\x8d\x73\x82\xa9\x82\xd3\x94\x4e\x82\xe0\x96\x94\x97\xb7\x90\x6c\x96\xe7\x81\x42", "月日は百代の過客にして、行かふ年も又旅人也。"},
	{UTF8, SJIS, "月日は百代の過客にして、行かふ年も又旅人也。", "\x8c\x8e\x93\xfa\x82\xcd\x95\x53\x91\xe3\x82\xcc\x89\xdf\x8b\x71\x82\xc9\x82\xb5\x82\xc4\x81\x41\x8d\x73\x82\xa9\x82\xd3\x94\x4e\x82\xe0\x96\x94\x97\xb7\x90\x6c\x96\xe7\x81\x42"},
	{EUCJP, UTF8, "\xa4\xb3\xa4\xf3\xa4\xcb\xa4\xc1\xa4\xcf\xa1\xa2Python\xa5\xd7\xa5\xed\xa5\xb0\xa5\xe9\xa5\xdf\xa5\xf3\xa5\xb0", "こんにちは、Pythonプログラミング"},
	{UTF8, EUCJP, "こんにちは、Pythonプログラミング", "\xa4\xb3\xa4\xf3\xa4\xcb\xa4\xc1\xa4\xcf\xa1\xa2Python\xa5\xd7\xa5\xed\xa5\xb0\xa5\xe9\xa5\xdf\xa5\xf3\xa5\xb0"},
}

var definiteTestData = []struct {
	BeforeCode, AfterCode CharCode
	BeforeText, AfterText string
}{
	{SJIS, UTF8, "\x82\xb1\x82\xea\x82\xcd\x8a\xbf\x8e\x9a\x82\xc5\x82\xb7\x81B", "これは漢字です。"},
	{UTF8, SJIS, "これは漢字です。", "\x82\xb1\x82\xea\x82\xcd\x8a\xbf\x8e\x9a\x82\xc5\x82\xb7\x81B"},
	{EUCJP, UTF8, "\xa4\xb3\xa4\xec\xa4\u03f4\xc1\xbb\xfa\xa4\u01e4\xb9\xa1\xa3", "これは漢字です。"},
	{SJIS, EUCJP, "\x82\xb1\x82\xea\x82\xcd\x8a\xbf\x8e\x9a\x82\xc5\x82\xb7\x81B", "\xa4\xb3\xa4\xec\xa4\u03f4\xc1\xbb\xfa\xa4\u01e4\xb9\xa1\xa3"},
	{EUCJP, SJIS, "\xa4\xb3\xa4\xec\xa4\u03f4\xc1\xbb\xfa\xa4\u01e4\xb9\xa1\xa3", "\x82\xb1\x82\xea\x82\xcd\x8a\xbf\x8e\x9a\x82\xc5\x82\xb7\x81B"},
	{UTF8, EUCJP, "これは漢字です。", "\xa4\xb3\xa4\xec\xa4\u03f4\xc1\xbb\xfa\xa4\u01e4\xb9\xa1\xa3"},
	{GBK, UTF8, "Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed", "Hello 常用國字標準字體表"},
	{UTF8, GBK, "Hello 常用國字標準字體表", "Hello \xb3\xa3\xd3\xc3\x87\xf8\xd7\xd6\x98\xcb\x9c\xca\xd7\xd6\xf3\x77\xb1\xed"},
	{UTF16, UTF8, "S0\x8c0o0\"oW[g0Y0\x020", "これは漢字です。"},
	{UTF8, UTF16, "これは漢字です。", "S0\x8c0o0\"oW[g0Y0\x020"},
	{UTF16, UTF8, "0S0\x8c0oo\"[W0g0Y0\x02", "これは漢字です。"},
	{UTF8, UTF16, "これは漢字です。", "0S0\x8c0oo\"[W0g0Y0\x02"},
}

// Test
func TestConvert(t *testing.T) {
	for _, data := range estimationTestData {
		converter := New(data.BeforeText, data.AfterCode)
		result, err := converter.Convert()
		if err != nil {
			panic(err)
		}
		assert.Equal(t, result, data.AfterText, "after convert")
	}

	for _, data := range definiteTestData {
		converter := New(data.BeforeText, data.AfterCode, data.BeforeCode)
		result, err := converter.Convert()
		if err != nil {
			panic(err)
		}
		assert.Equal(t, result, data.AfterText, "after convert")
	}
}

func benchmarkConvert(b *testing.B, text string, codes ...CharCode) {
	for i := 0; i < b.N; i++ {
		converter := New(text, codes...)
		converter.Convert()
	}
}

// Benchmark
func BenchmarkEstimation(b *testing.B) {
	benchmarkConvert(b, estimationTestData[0].BeforeText, estimationTestData[0].AfterCode)
}
func BenchmarkConvertWithoutEstimation(b *testing.B) {
	benchmarkConvert(b, definiteTestData[0].BeforeText, definiteTestData[0].AfterCode)
}
func BenchmarkConvertWithMediation(b *testing.B) {
	benchmarkConvert(b, definiteTestData[3].BeforeText, definiteTestData[3].AfterCode)
}
