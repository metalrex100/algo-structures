package algo_structures

type vectorArray struct {
	*baseArray
	resizeStep int
}

func NewVectorArray(resizeStep int) *vectorArray {
	return &vectorArray{
		baseArray:  newBaseArray(),
		resizeStep: resizeStep,
	}
}

func (a *vectorArray) Add(item int) {
	if a.isArrayFull() {
		a.resize(len(a.arr) + a.resizeStep)
	}

	a.arr[a.cursor] = item
	a.cursor++
}
