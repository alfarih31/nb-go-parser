package parser

import (
	"database/sql/driver"
	"fmt"
)

func toInt16Ptr(i int16) *int16 {
	return &i
}

func toIntPtr(i int) *int {
	return &i
}

func toInt32Ptr(i int32) *int32 {
	return &i
}

func toInt64Ptr(i int64) *int64 {
	return &i
}

type IntPtr struct {
	i *int
}

type Int16Ptr struct {
	i *int16
}

type Int32Ptr struct {
	i *int32
}

type Int64Ptr struct {
	i *int64
}

func (p IntPtr) Value() (driver.Value, error) {
	return p.i, nil
}

func (p *IntPtr) Scan(value interface{}) error {
	if value == nil {
		p.i = nil
		return nil
	}

	if bv, err := driver.Int32.ConvertValue(value); err != nil {
		if v, ok := bv.(int); ok {
			*p = IntPtr{
				i: toIntPtr(v),
			}
		}
	} else {
		return fmt.Errorf("failed to scan IntPtr:\n%v", err)
	}

	return nil
}

func (p Int32Ptr) Value() (driver.Value, error) {
	return p.i, nil
}

func (p *Int32Ptr) Scan(value interface{}) error {
	if value == nil {
		p.i = nil
		return nil
	}

	if bv, err := driver.Int32.ConvertValue(value); err != nil {
		if v, ok := bv.(int32); ok {
			*p = Int32Ptr{
				i: toInt32Ptr(v),
			}
		}
	} else {
		return fmt.Errorf("failed to scan Int32Ptr:\n%v", err)
	}

	return nil
}

func (p Int64Ptr) Value() (driver.Value, error) {
	return p.i, nil
}

func (p *Int64Ptr) Scan(value interface{}) error {
	if value == nil {
		p.i = nil
		return nil
	}

	if bv, err := driver.Int32.ConvertValue(value); err != nil {
		if v, ok := bv.(int64); ok {
			*p = Int64Ptr{
				i: toInt64Ptr(v),
			}
		}
	} else {
		return fmt.Errorf("failed to scan Int64Ptr:\n%v", err)
	}

	return nil
}

func (p Int16Ptr) Value() (driver.Value, error) {
	return p.i, nil
}

func (p *Int16Ptr) Scan(value interface{}) error {
	if value == nil {
		p.i = nil
		return nil
	}

	if bv, err := driver.Int32.ConvertValue(value); err != nil {
		if v, ok := bv.(int16); ok {
			*p = Int16Ptr{
				i: toInt16Ptr(v),
			}
		}
	} else {
		return fmt.Errorf("failed to scan Int16Ptr:\n%v", err)
	}

	return nil
}
