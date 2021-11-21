package algo_structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorArray(t *testing.T) {
	a := NewFactorArray()

	ast := assert.New(t)
	ast.Equal(0, a.Count())

	a.Add(7)

	ast.Equal(1, a.Count())
	ast.Equal(7, a.Get(0))

	a.Add(12)

	ast.Equal(2, a.Count())
	ast.Equal(12, a.Get(1))

	a.Add(15)

	ast.Equal(4, a.Count())
	ast.Equal(15, a.Get(2))

	a.Add(20)

	ast.Equal(4, a.Count())
	ast.Equal(20, a.Get(3))

	a.Add(22)

	ast.Equal(8, a.Count())
	ast.Equal(22, a.Get(4))

	a.RemoveLast()

	ast.Equal(7, a.Count())
	ast.Panics(func() {
		a.Get(7)
	})
}