package parser

type Bool bool

func (b Bool) ToString() string {
	if b {
		return "true"
	}

	return "false"
}

func (b Bool) ToInt() int {
	if b {
		return 1
	}

	return 0
}
