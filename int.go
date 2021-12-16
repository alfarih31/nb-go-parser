package parser

import (
	"strconv"
)

type Int int

func (i Int) ToString() (string, error) {
	return strconv.Itoa(int(i)), nil
}
