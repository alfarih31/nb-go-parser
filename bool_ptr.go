package parser

import (
	"database/sql/driver"
	"fmt"
)

func toBoolPtr(b bool) *bool {
	return &b
}

type BoolPtr struct {
	b *bool
}

func (p BoolPtr) Value() (driver.Value, error) {
	return p.b, nil
}

func (p *BoolPtr) Scan(value interface{}) error {
	if value == nil {
		p.b = nil
		return nil
	}

	if bv, err := driver.Bool.ConvertValue(value); err != nil {
		if v, ok := bv.(bool); ok {
			*p = BoolPtr{
				b: toBoolPtr(v),
			}
		}
	} else {
		return fmt.Errorf("failed to scan BoolPtr:\n%v", err)
	}

	return nil
}
