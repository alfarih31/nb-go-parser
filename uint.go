package parser

import "strconv"

type UInt uint

// ToString cast UInt to string
func (u UInt) ToString() (string, error) {
	return strconv.Itoa(int(u)), nil
}

// ToBool cast UInt to bool. Returns true for i > 0 & return false for i <= 0
func (u UInt) ToBool() (bool, error) {
	if u > 0 {
		return true, nil
	}
	return false, nil
}

func (i UInt) ToUInt() uint {
	return uint(i)
}

func (i UInt) ToUInt32() uint32 {
	return uint32(i)
}

func (i UInt) ToUInt64() uint64 {
	return uint64(i)
}

func (i UInt) ToUIntPtr() *uint {
	return UIntPtr(uint(i))
}

func (i UInt) ToUIntPtr32() *uint32 {
	return UIntPtr32(i.ToUInt32())
}

func (i UInt) ToUIntPtr64() *uint64 {
	return UIntPtr64(i.ToUInt64())
}

func UIntPtr(i uint) *uint {
	return &i
}

func UIntPtr32(i uint32) *uint32 {
	return &i
}

func UIntPtr64(i uint64) *uint64 {
	return &i
}
