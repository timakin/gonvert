package converter

import (
	"bytes"
	"golang.org/x/text/encoding/unicode"
)

type UTF8ToUTF16BEConverter ConversionPattern

func (c *UTF8ToUTF16BEConverter) Convert() (string, error) {
	return transformEncoding(bytes.NewReader(c.TextByte), unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewEncoder())
}

type UTF8ToUTF16LEConverter ConversionPattern

func (c *UTF8ToUTF16LEConverter) Convert() (string, error) {
	return transformEncoding(bytes.NewReader(c.TextByte), unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder())
}
