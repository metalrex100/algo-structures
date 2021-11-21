package algo_structures

type FactorArray struct {
	*baseArray
}

func NewFactorArray() *FactorArray {
	return &FactorArray{baseArray: newBaseArray()}
}

func (a *FactorArray) Add(item int) {
	if a.isArrayFull() {
		newSize := len(a.arr) * 2
		if newSize < 1 {
			newSize = 1
		}
		a.resize(newSize)
	}
	a.arr[a.cursor] = item
	a.cursor++
}
