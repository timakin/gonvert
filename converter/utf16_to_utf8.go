package converter

import (
	"bytes"
	"unicode/utf16"
	"unicode/utf8"
)

type UTF16ToUTF8Converter ConversionPattern

func (c *UTF16ToUTF8Converter) Convert() (string, error) {
	u16s := make([]uint16, 1)

	ret := &bytes.Buffer{}

	b8buf := make([]byte, 4)
	b := c.TextByte
	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}
