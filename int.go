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

func (i Int) ToInt() int {
	return int(i)
}

func (i Int) ToInt32() int32 {
	return int32(i)
}

func (i Int) ToInt64() int64 {
	return int64(i)
}

func (i Int) ToIntPtr() *int {
	return toIntPtr(int(i))
}

func (i Int) ToInt32Ptr() *int32 {
	return toInt32Ptr(i.ToInt32())
}

func (i Int) ToInt64Ptr() *int64 {
	return toInt64Ptr(i.ToInt64())
}
