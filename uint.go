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

func (u UInt) ToUInt() uint {
	return uint(u)
}

func (u UInt) ToUInt32() uint32 {
	return uint32(u)
}

func (u UInt) ToUInt64() uint64 {
	return uint64(u)
}

func (u UInt) ToUIntPtr() *uint {
	return UIntPtr(uint(u))
}

func (u UInt) ToUIntPtr32() *uint32 {
	return UIntPtr32(u.ToUInt32())
}

func (u UInt) ToUIntPtr64() *uint64 {
	return UIntPtr64(u.ToUInt64())
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
