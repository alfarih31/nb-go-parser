package parser

type IntArr []int

func (a IntArr) ToInt32Arr() []int32 {
	o := make([]int32, len(a), cap(a))

	for idx, i := range a {
		o[idx] = Int(i).ToInt32()
	}

	return o
}

func (a IntArr) ToInt64Arr() []int64 {
	o := make([]int64, len(a), cap(a))

	for idx, i := range a {
		o[idx] = Int(i).ToInt64()
	}

	return o
}

func (a IntArr) ToInt32PtrArr() []*int32 {
	o := make([]*int32, len(a), cap(a))

	for idx, i := range a {
		o[idx] = Int(i).ToIntPtr32()
	}

	return o
}

func (a IntArr) ToInt64PtrArr() []*int64 {
	o := make([]*int64, len(a), cap(a))

	for idx, i := range a {
		o[idx] = Int(i).ToIntPtr64()
	}

	return o
}

type Int32Arr []int32

func (a Int32Arr) ToIntArr() []int {
	o := make([]int, len(a), cap(a))

	for idx, i := range a {
		o[idx] = Int(i).ToInt()
	}

	return o
}

type Int64Arr []int64

func (a Int64Arr) ToIntArr() []int {
	o := make([]int, len(a), cap(a))

	for idx, i := range a {
		o[idx] = Int(i).ToInt()
	}

	return o
}
