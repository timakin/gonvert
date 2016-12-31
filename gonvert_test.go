package gonvert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	converter := New("\x89\xD4\x8e\x71", UTF8)
	result, err := converter.Convert()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, result, "花子", "they should be equal")
}
