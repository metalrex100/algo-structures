package algo_structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorArray(t *testing.T) {
	a := NewVectorArray(3)

	ast := assert.New(t)
	ast.Equal(0, a.Count())

	a.Add(7)
	ast.Equal(3, a.Count())
	ast.Equal(7, a.Get(0))

	a.Add(12)
	a.Add(15)

	ast.Equal(3, a.Count())
	ast.Equal(15, a.Get(2))

	a.Add(20)

	ast.Equal(6, a.Count())
	ast.Equal(20, a.Get(3))

	a.RemoveLast()

	ast.Equal(5, a.Count())
	ast.Panics(func() {
		a.Get(5)
	})
}
