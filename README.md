Gonvert
====

Dead simple character-encoding converter in Golang.
You can convert without a declaration of `before` encoding.

## Install
```
go get github.com/timakin/gonvert
```

## Support
- Shift_JIS -> UTF8
- UTF8 -> Shift_JIS

## Usage

```
package main

import (
    "github.com/timakin/gonvert"
    "fmt"
)
func main() {
    converter := gonvert.New()
    result, err := converter.Convert(/* string */, gonvert.UTF8)
    if err != nil {
        panic("Failed to convert!")
    }
    fmt.Print(result) // Output is utf-8 encoded string
}
```

