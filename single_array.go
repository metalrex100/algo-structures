package algo_structures

type singleArray struct {
	*baseArray
}

func NewSingleArray() *singleArray {
	return &singleArray{newBaseArray()}
}

func (a *singleArray) Count() int {
	return len(a.arr)
}

func (a *singleArray) Add(item int) {
	if a.isArrayFull() {
		a.resize(len(a.arr) + 1)
	}

	a.arr[a.cursor] = item
	a.cursor++
}
