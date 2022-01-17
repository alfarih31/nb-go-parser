package parser

type Bool bool

// ToString cast bool to represented string
func (b Bool) ToString() string {
	if b {
		return "true"
	}

	return "false"
}

// ToInt cast bool to represented bool
func (b Bool) ToInt() int {
	if b {
		return 1
	}

	return 0
}

func (b Bool) ToBoolPtr() *bool {
	return toBoolPtr(bool(b))
}
