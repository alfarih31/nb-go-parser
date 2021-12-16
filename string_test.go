package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString_ToBool(t *testing.T) {
	b, err := String("false").ToBool()
	assert.Equal(t, err, nil, "Error must be nil")
	assert.Equal(t, b, false, "Expected output is: false")
}

func TestString_ToInt(t *testing.T) {
	i, err := String("1234").ToInt()
	assert.Equal(t, err, nil, "Error must be nil")
	assert.Equal(t, i, 1234, "Expected output is: 1234")
}

func TestString_ToStringArr(t *testing.T) {
	sr, err := String("first:second:third").ToStringArr(":")
	assert.Equal(t, err, nil, "Error must be nil")
	assert.Equal(t, sr, []string{"first", "second", "third"}, `Expected output is: ["first", "second", "third"]`)
}

func TestString_ToIntArr(t *testing.T) {
	ir, err := String("1:2:3").ToIntArr(":")
	assert.Equal(t, err, nil, "Error must be nil")
	assert.Equal(t, ir, []int{1, 2, 3}, `Expected output is: [1, 2, 3]`)
}

func TestString_MaskRight(t *testing.T) {
	ts := String("NOT_MASKED:MASKED").MaskRight(11, '*')

	assert.Equal(t, ts, "NOT_MASKED:******", "Expected output is: NOT_MASKED:******")
}

func TestString_MaskLeft(t *testing.T) {
	ts := String("MASKED:NOT_MASKED").MaskLeft(11, '*')

	assert.Equal(t, ts, "******:NOT_MASKED", "Expected output is: ******:NOT_MASKED")
}
