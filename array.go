package algo_structures

type Array interface {
	Count() int
	Add(item int)
	RemoveLast()
	Get(index int) int
}

type baseArray struct {
	arr    []int
	cursor int
}

func newBaseArray() *baseArray {
	return &baseArray{arr: make([]int, 0)}
}

func (a *baseArray) Count() int {
	return len(a.arr)
}

func (a *baseArray) RemoveLast() {
	if len(a.arr) < 1 {
		return
	}

	a.resize(len(a.arr) - 1)
	if a.cursor > a.getLastElemIndex() {
		a.cursor = a.getLastElemIndex()
	}
}

func (a *baseArray) Get(index int) int {
	return a.arr[index]
}

func (a *baseArray) resize(newSize int) {
	newArr := make([]int, newSize)
	lastElemIndex := a.getLastElemIndex()

	for i := 0; i < lastElemIndex; i++ {
		newArr[i] = a.arr[i]
	}
	a.arr = newArr
}

func (a *baseArray) isArrayFull() bool {
	return a.cursor > a.getLastElemIndex()
}

func (a *baseArray) getLastElemIndex() int {
	return len(a.arr) - 1
}
