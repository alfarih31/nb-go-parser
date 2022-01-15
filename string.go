package parser

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type String string

// ToInt cast String to int
func (s String) ToInt() (int, error) {
	return strconv.Atoi(string(s))
}

func (s String) ToStringPtr() *string {
	return toStringPtr(string(s))
}

// ToBool cast String to represented bool
func (s String) ToBool() (bool, error) {
	switch string(s) {
	case "true":
		fallthrough
	case "True":
		fallthrough
	case "TRUE":
		fallthrough
	case "1":
		return true, nil
	case "false":
		fallthrough
	case "False":
		fallthrough
	case "FALSE":
		fallthrough
	case "0":
		return false, nil
	default:
		return false, errors.New("string cannot be converted to bool")
	}
}

// ToStringArr cast String to array of string
func (s String) ToStringArr(separator ...string) ([]string, error) {
	sep := ","
	if len(separator) > 0 {
		sep = separator[0]
	}

	var ss []string
	for _, v := range strings.Split(string(s), sep) {
		ss = append(ss, strings.TrimSpace(v))
	}

	return ss, nil
}

// ToIntArr cast String to array of int
func (s String) ToIntArr(separator ...string) ([]int, error) {
	sep := ","
	if len(separator) > 0 {
		sep = separator[0]
	}

	var is []int
	for _, v := range strings.Split(string(s), sep) {
		i, e := strconv.Atoi(strings.TrimSpace(v))

		if e != nil {
			return nil, e
		}

		is = append(is, i)
	}

	return is, nil
}

// MaskRight mask string with a char(rune) and left n value on the right side
// Example: String("NOT_MASKED:MASKED").MaskRight(11, '*') returns NOT_MASKED:******
func (s String) MaskRight(NLeft int, char rune) string {
	rs := []rune(s)

	n := len(s)
	// Bound NLeft to max length of text
	maxNLeft := int(math.Max(math.Min(float64(n), float64(NLeft)), 0))

	for i := 0; i < n-maxNLeft; i++ {
		rs[n-i-1] = char
	}

	return string(rs)
}

// MaskLeft mask string with a char(rune) and left n value on the left side
// Example: String("MASKED:NOT_MASKED").MaskLeft(11, '*') returns ******:NOT_MASKED
func (s String) MaskLeft(NLeft int, char rune) string {
	rs := []rune(s)

	n := len(s)
	// Bound NLeft to max length of text
	maxNLeft := int(math.Max(math.Min(float64(n), float64(NLeft)), 0))

	for i := 0; i < n-maxNLeft; i++ {
		rs[i] = char
	}

	return string(rs)
}
