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

func NilAbleUInt(u uint) *uint {
	return &u
}
