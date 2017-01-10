package converter

import (
	"bytes"
	"golang.org/x/text/encoding/unicode"
)

type UTF16BEToUTF8Converter ConversionPattern

func (c *UTF16BEToUTF8Converter) Convert() (string, error) {
	return transformEncoding(bytes.NewReader(c.TextByte), unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewDecoder())
}

type UTF16LEToUTF8Converter ConversionPattern

func (c *UTF16LEToUTF8Converter) Convert() (string, error) {
	return transformEncoding(bytes.NewReader(c.TextByte), unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder())
}
