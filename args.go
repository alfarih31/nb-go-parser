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

func GetOptInt32Arg(opts []int32, defaultValue ...int32) int32 {
	var o int32 = 0
	if len(defaultValue) > 0 {
		o = defaultValue[0]
	}

	if len(opts) > 0 {
		o = opts[0]
	}

	return o
}

func GetOptInt64Arg(opts []int64, defaultValue ...int64) int64 {
	var o int64 = 0
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

	if len(opts) > 0 && opts[0] != "" {
		o = opts[0]
	}

	return o
}
