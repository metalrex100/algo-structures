package algo_structures

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedArray(t *testing.T) {
	ast := assert.New(t)

	la := NewLinkedList()

	ast.Equal(0, la.Len())

	la.AddToTheEnd(5)
	ast.Equal(1, la.Len())
	ast.Equal(la.GetByIndex(0), 5)

	la.Print()

	la.AddToTheEnd(7)
	ast.Equal(2, la.Len())
	ast.Equal(la.GetByIndex(1), 7)

	la.Print()

	la.AddToTheEnd(9)
	ast.Equal(3, la.Len())
	ast.Equal(la.GetByIndex(2), 9)

	la.Print()

	la.InsertByIndex(1, 14)
	ast.Equal(4, la.Len())
	ast.Equal(14, la.GetByIndex(1))

	la.Print()

	la.InsertByIndex(9, 34)
	ast.Equal(10, la.Len())
	ast.Equal(34, la.GetByIndex(9))

	la.Print()

	la.InsertByIndex(3, 22)
	ast.Equal(11, la.Len())
	ast.Equal(22, la.GetByIndex(3))

	la.Print()

	la.RemoveLast()
	ast.Equal(10, la.Len())
	ast.Panics(func() {
		la.GetByIndex(10)
	})

	la.Print()

	la.RemoveByIndex(3)
	ast.Equal(9, la.Len())
	ast.Equal(9, la.GetByIndex(3))

	la.Print()

	la.RemoveByIndex(1)
	ast.Equal(8, la.Len())
	ast.Equal(7, la.GetByIndex(1))

	la.Print()

	la.RemoveByIndex(0)
	ast.Equal(7, la.Len())
	ast.Equal(7, la.GetByIndex(0))

	la.Print()
}
