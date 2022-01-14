package parser

import (
	"database/sql/driver"
	"fmt"
)

type StringPtr struct {
	s *string
}

func toStringPtr(s string) *string {
	return &s
}

func (p StringPtr) Value() (driver.Value, error) {
	return p.s, nil
}

func (p *StringPtr) Scan(value interface{}) error {
	if value == nil {
		p.s = nil
		return nil
	}

	if bv, err := driver.String.ConvertValue(value); err != nil {
		if v, ok := bv.(string); ok {
			*p = StringPtr{
				s: toStringPtr(v),
			}
		}
	} else {
		return fmt.Errorf("failed to scan StringPtr:\n%v", err)
	}

	return nil
}
