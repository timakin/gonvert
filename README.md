Gonvert
====

Dead simple character-encoding converter in Golang.
You can convert without a declaration of `before` encoding.

## Install
```
go get github.com/timakin/gonvert
```

## Current Support
- Shift_JIS <-> UTF8
- EUC-JP <-> UTF8

## Usage

```
package main

import (
    "github.com/timakin/gonvert"
    "fmt"
)
func main() {
    // Input a Shift_JIS encoded string
    sjisStr := "\x8c\x8e\x93\xfa\x82\xcd\x95\x53\x91\xe3\x82\xcc\x89\xdf\x8b" +
               "\x71\x82\xc9\x82\xb5\x82\xc4\x81\x41\x8d\x73\x82\xa9\x82\xd3" +
               "\x94\x4e\x82\xe0\x96\x94\x97\xb7\x90\x6c\x96\xe7\x81\x42"
    converter := gonvert.New(sjisStr, gonvert.UTF8)
    result, err := converter.Convert()
    if err != nil {
        panic("Failed to convert!")
    }
    fmt.Print(result) // It will print out the utf-8 encoded string: "月日は百代の過客にして、行かふ年も又旅人也。"
}
```
