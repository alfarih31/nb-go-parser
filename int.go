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
	return IntPtr(int(i))
}

func (i Int) ToIntPtr32() *int32 {
	return IntPtr32(i.ToInt32())
}

func (i Int) ToIntPtr64() *int64 {
	return IntPtr64(i.ToInt64())
}

func IntPtr(i int) *int {
	return &i
}

func IntPtr32(i int32) *int32 {
	return &i
}

func IntPtr64(i int64) *int64 {
	return &i
}
