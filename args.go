package parser

func GetOptArg(opts []interface{}, defaultValue ...interface{}) interface{} {
	var o interface{} = nil
	if len(defaultValue) > 0 {
		o = defaultValue[0]
	}

	if len(opts) > 0 {
		o = opts[0]
	}

	return o
}

func GetOptIntArg(opts []int, defaultValue ...int) int {
	var o = 0
	if len(defaultValue) > 0 {
		o = defaultValue[0]
	}

	if len(opts) > 0 {
		o = opts[0]
	}

	return o
}

func GetOptBoolArg(opts []bool, defaultValue ...bool) bool {
	var o = false
	if len(defaultValue) > 0 {
		o = defaultValue[0]
	}

	if len(opts) > 0 {
		o = opts[0]
	}

	return o
}

func GetOptStringArg(opts []string, defaultValue ...string) string {
	var o = ""
	if len(defaultValue) > 0 {
		o = defaultValue[0]
	}

	if len(opts) > 0 {
		o = opts[0]
	}

	return o
}
