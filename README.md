Gonvert
====

Simple character-encoding converter with an automatic character-code detection in Golang.
You can convert without a declaration of `before` encoding.

## Install
```
go get github.com/timakin/gonvert
```

## Current Support
- Shift_JIS <-> UTF8
- EUC-JP <-> UTF8
- GBK <-> UTF8

You can specify the character code to encode/decode with gonvert constatants.

Prepared `const` character-code is following.

```
const (
	UTF8 CharCode = iota
	SJIS
	EUCJP
	GBK
)
```

## Usage

You can call the converter with 2 or 3 arguements.

If you set 2 variables, gonvert will estimate the code automatically.

But if you already know the code of strings, you should set the third arguements, without an estimation.

```
package main

import (
    "github.com/timakin/gonvert"
    "fmt"
)
func main() {
    // ------------ Estimation case ------------

    // Input a Shift_JIS encoded string
    sjisStr := "\x8c\x8e\x93\xfa\x82\xcd\x95\x53\x91\xe3\x82\xcc\x89\xdf\x8b" +
               "\x71\x82\xc9\x82\xb5\x82\xc4\x81\x41\x8d\x73\x82\xa9\x82\xd3" +
               "\x94\x4e\x82\xe0\x96\x94\x97\xb7\x90\x6c\x96\xe7\x81\x42"
    converter := gonvert.New(sjisStr, gonvert.UTF8)
    result, err := converter.Convert()
    if err != nil {
        panic("Failed to convert!")
    }
    // This will print out the utf-8 encoded string: "月日は百代の過客にして、行かふ年も又旅人也。"
    fmt.Print(result)

    // -----------------------------------------

    // ------------ Specified-code case ------------

    sjisStr := "\x8c\x8e\x93\xfa\x82\xcd\x95\x53\x91\xe3\x82\xcc\x89\xdf\x8b" +
               "\x71\x82\xc9\x82\xb5\x82\xc4\x81\x41\x8d\x73\x82\xa9\x82\xd3" +
               "\x94\x4e\x82\xe0\x96\x94\x97\xb7\x90\x6c\x96\xe7\x81\x42"

    // Should send `before` character code if you already know, because an estimation like above may become incorrect.
    converter := gonvert.New(sjisStr, gonvert.UTF8, gonvert.SJIS)
    result, err := converter.Convert()
    if err != nil {
        panic("Failed to convert!")
    }
    fmt.Print(result)

    // -----------------------------------------
}
```
