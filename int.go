package parser

import (
	"strconv"
)

type Int int

// ToString cast Int to string
func (i Int) ToString() (string, error) {
	return strconv.Itoa(int(i)), nil
}

// ToBool cast Int to bool. Returns true for i > 0 & return false for i <= 0
func (i Int) ToBool() (bool, error) {
	if i > 0 {
		return true, nil
	}
	return false, nil
}
