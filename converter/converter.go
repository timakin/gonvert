package converter

type Converter interface {
	Convert() (string, error)
}

type ConversionPattern struct {
	TextByte []byte
}

// func transformEncoding(rawReader io.Reader, trans transform.Transformer) (string, error) {
// 	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
// 	if err == nil {
// 		return string(ret), nil
// 	} else {
// 		return "", err
// 	}
// }
//
// // Convert a string encoding from ShiftJIS to UTF-8
// func FromShiftJIS(str string) (string, error) {
// 	return transformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewDecoder())
// }
