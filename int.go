package parser

import (
	"strconv"
)

type Int int

// ToString cast Int to string
func (i Int) ToString() string {
	return strconv.Itoa(int(i))
}

// ToBool cast Int to bool. Returns true for i > 0 & return false for i <= 0
func (i Int) ToBool() bool {
	if i > 0 {
		return true
	}
	return false
}

func (i Int) ToInt() int {
	return int(i)
}

func (i Int) ToInt16() int16 {
	return int16(i)
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

func (i Int) ToInt16Ptr() *int16 {
	return toInt16Ptr(int16(i))
}

func (i Int) ToInt32Ptr() *int32 {
	return toInt32Ptr(i.ToInt32())
}

func (i Int) ToInt64Ptr() *int64 {
	return toInt64Ptr(i.ToInt64())
}
